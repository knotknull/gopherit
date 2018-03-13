package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/hashicorp/memberlist"
)

func main() {
	list, err := memberlist.Create(memberlist.DefaultLocalConfig())
	if err != nil {
		panic("Failed to create memberlist: " + err.Error())
	}

	// Join an existing cluster by specifying at least one known member.
	_, err = list.Join([]string{"0.0.0.0"})
	if err != nil {
		panic("Failed to join cluster: " + err.Error())
	}

	// Ask for members of the cluster
	for _, member := range list.Members() {
		fmt.Printf("Member: %s %s\n", member.Name, member.Addr)
	}

	// Wire up an interrupt so we don't exit early
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	// Block until a signal is received.
	s := <-c
	fmt.Println("Got signal:", s)
}
