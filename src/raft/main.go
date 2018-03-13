package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gopherguides/training/distributed-systems/consensus/src/raft/httpd"
	"github.com/gopherguides/training/distributed-systems/consensus/src/raft/store"
)

const (
	defaultHTTPAddr = ":8080"
	defaultRaftAddr = ":8086"
)

// Command line parameters
var httpAddr string
var raftAddr string
var joinAddr string

// bindFlags will bind the command arguments to our variables
func bindFlags() {
	// Read and bind all
	flag.StringVar(&httpAddr, "httpaddr", defaultHTTPAddr, "Set the HTTP bind address")
	flag.StringVar(&raftAddr, "raftaddr", defaultRaftAddr, "Set Raft bind address")
	flag.StringVar(&joinAddr, "join", "", "[optional] The address of a node to join.  Leave empty to boostrap your first node.")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <raft-data-path> \n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Usage = usage
}

func main() {
	bindFlags()
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "No Raft storage directory specified\n")
		os.Exit(1)
	}

	// Ensure Raft storage exists.
	raftDir := flag.Arg(0)
	if raftDir == "" {
		fmt.Fprintf(os.Stderr, "No Raft storage directory specified\n")
		os.Exit(1)
	}
	os.MkdirAll(raftDir, 0700)

	// Create a new store
	s := store.New()
	if err := s.Open(joinAddr, raftDir, raftAddr); err != nil {
		log.Fatalf("failed to open store: %s", err.Error())
	}

	// Create our new httpd server. Pass it a reference of the raft store
	h := httpd.New(httpAddr, s)
	if err := h.Start(); err != nil {
		log.Fatalf("failed to start HTTP service: %s", err.Error())
	}
	log.Printf("started http service on %s", httpAddr)

	// If join was specified, make the join request.
	if joinAddr != "" {
		if err := join(joinAddr, raftAddr); err != nil {
			log.Fatalf("failed to join node at %s: %s", joinAddr, err.Error())
		}
	}

	// wire up an interrupt so we don't exit until we are told to do so.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt
	log.Println("interrupt signal received, terminating")
}

// Since you can only add a peer if you are the leader
// we need to actually post a request to this address
// and request we join
func join(joinAddr, raftAddr string) error {
	b, err := json.Marshal(map[string]string{"addr": raftAddr})
	if err != nil {
		return err
	}
	host := fmt.Sprintf("http://%s/join", joinAddr)

	log.Printf("attempting to join %s", host)
	resp, err := http.Post(host, "application-type/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func usage() {
	fmt.Println(`
usage: raft [options] path

	path: path on disk to store your raft database and peers.json

Start your first node:

raft -httpaddr localhost:8180 --raftaddr localhost:8186 /tmp/raft1

Start your second node:

	raft --httpaddr localhost:8280 --raftaddr localhost:8286 --join localhost:8180 /tmp/raft2

Start your third node:

	raft --httpaddr localhost:8380 --raftaddr localhost:8386 --join localhost:8180 /tmp/raft3


Options:
	`)
	flag.PrintDefaults()
}
