package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gopherguides/training/distributed-systems/grpc/src/cachely/cachely"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	// connect to the grpc server
	conn, err := grpc.Dial(":5051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	// create a new client
	c := cachely.NewCacheClient(conn)

	ctx := context.Background()

	// write a value
	_, err = c.Put(ctx, &cachely.PutRequest{
		Key:   "band",
		Value: []byte("The Beatles"),
	})
	if err != nil {
		log.Fatal(err)
	}

	// get a value
	gr, err := c.Get(ctx, &cachely.GetRequest{Key: "band"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("band:", string(gr.GetValue()))

	// delete a value
	_, err = c.Delete(ctx, &cachely.DeleteRequest{Key: "band"})
	if err != nil {
		log.Fatal(err)
	}

	// check it was deleted
	gr, err = c.Get(ctx, &cachely.GetRequest{Key: "band"})
	code := status.Code(err)
	switch code {
	case codes.NotFound:
		fmt.Println("confirmed 'band' has been deleted")
	case codes.OK:
		log.Fatal("band code was found, delete failed")
	default:
		log.Fatal(err)
	}
}
