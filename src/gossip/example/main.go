package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/hashicorp/memberlist"
	uuid "github.com/satori/go.uuid"
)

var (
	broadcasts *memberlist.TransmitLimitedQueue
)

func main() {

	// Read in our flags
	members := flag.String("members", "", "comma seperated list of members")
	mport := flag.Int("mport", 0, "memberlist port (0 = auto select)")
	port := flag.Int("port", 4001, "http port")

	// parse the flags
	flag.Parse()

	// Create our cache (in memory concurrency safe)
	cache := NewCache()

	// Create a delegate
	d := NewDelegate(cache)

	// We'll use a default config for Memberlist.  Their are several `sane`
	// configs in the library to use for development and testing.
	// reference: https://godoc.org/github.com/hashicorp/memberlist#Config
	c := memberlist.DefaultWANConfig()

	// Set the bind port for the service to use
	c.BindPort = *mport

	// Set our delegate
	c.Delegate = d

	// Create a unique host name for memberlist
	hostname, _ := os.Hostname()

	// We want to identify each node with a unique identity.  For simplicity, we'll
	// use the host name and generate a UUID.
	if uid, err := uuid.NewV4(); err != nil {
		panic(err)
	} else {
		c.Name = hostname + "-" + uid.String()
	}

	// Create will create a new Memberlist using the given configuration. This
	// will not connect to any other node (see Join) yet, but will start all the
	// listeners to allow other nodes to join this memberlist. After creating a
	// Memberlist, the configuration given should not be modified by the user
	// anymore.
	m, err := memberlist.Create(c)
	if err != nil {
		if err != nil {
			panic("Failed to create memberlist: " + err.Error())
		}
	}

	// Join is used to take an existing Memberlist and attempt to join a cluster by
	// contacting all the given hosts and performing a state sync. Initially, the
	// Memberlist only contains our own state, so doing this will cause remote nodes
	// to become aware of the existence of this node, effectively joining the cluster.

	// This returns the number of hosts successfully contacted and an error if none
	// could be reached. If an error is returned, the node did not successfully join
	// the cluster.

	// If we passed in members, we need to join them
	if len(*members) > 0 {
		parts := strings.Split(*members, ",")
		n, err := m.Join(parts)
		if err != nil {
			panic(fmt.Sprintf("error joining %q: %s", *members, err))
		}
		log.Printf("%d members currently known\n", n)
	}

	// TransmitLimitedQueue is used to queue messages to broadcast to the cluster
	// (via gossip) but limits the number of transmits per message. It also
	// prioritizes messages with lower transmit counts (hence newer messages).
	// reference:
	// https://godoc.org/github.com/hashicorp/memberlist#TransmitLimitedQueue

	// Wire up the broadcasts
	broadcasts = &memberlist.TransmitLimitedQueue{
		NumNodes: func() int {
			return m.NumMembers()
		},
		RetransmitMult: 3,
	}

	// Get a reference to our local node, and print out the current state of our cluster.
	node := m.LocalNode()
	fmt.Printf("Local member %s:%d\n", node.Addr, node.Port)

	// Ask for members of the cluster and print them out. Members returns a list
	// of all known live nodes. The node structures returned must not be
	// modified. If you wish to modify a Node, make a copy first.
	for _, member := range m.Members() {
		fmt.Printf("Member: %s\nAddr: %s\nPort: %d\n", member.Name, member.Addr, member.Port)
	}

	// Wire up the api endpoints
	s := NewServer(cache)
	http.HandleFunc("/add", s.addHandler)
	http.HandleFunc("/del", s.delHandler)
	http.HandleFunc("/get", s.getHandler)
	fmt.Printf("Webserver Listening on :%d\n", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		fmt.Println(err)
	}
}
