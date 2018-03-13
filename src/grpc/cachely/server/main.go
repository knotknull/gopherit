package main

import (
	"log"
	"net"
	"sync"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/gopherguides/training/distributed-systems/grpc/src/cachely/cachely"
)

// Server is the main concept that will house our cache of values
type server struct {
	data sync.Map
}

// Get will retrieve a value from the map and return it if found
// Get returns an error if the key is not found.
func (s *server) Get(_ context.Context, req *cachely.GetRequest) (*cachely.GetResponse, error) {
	key := req.GetKey()
	log.Printf("looking up key %q\n", key)
	if v, ok := s.data.Load(key); ok {
		log.Printf("found key %q\n", key)
		return &cachely.GetResponse{
			Key:   key,
			Value: v.([]byte),
		}, nil
	}
	log.Printf("key not found %q\n", key)
	return nil, status.Errorf(codes.NotFound, "could not find key %s", key)
}

// Put will store a new value or update the existing value
func (s *server) Put(_ context.Context, req *cachely.PutRequest) (*cachely.PutResponse, error) {
	log.Printf("storing key %q\n", req.GetKey())
	s.data.Store(req.GetKey(), req.GetValue())
	return &cachely.PutResponse{
		Key: req.GetKey(),
	}, status.New(codes.OK, "").Err()
}

// Delete will remove the key.  It does not error out if the key is not found.
func (s *server) Delete(_ context.Context, req *cachely.DeleteRequest) (*cachely.DeleteResponse, error) {
	log.Printf("deleting key %q\n", req.GetKey())
	s.data.Delete(req.GetKey())
	return &cachely.DeleteResponse{
		Key: req.GetKey(),
	}, status.New(codes.OK, "").Err()
}

func main() {
	// open a port to communicate on
	lis, err := net.Listen("tcp", ":5051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a new grpc server
	s := grpc.NewServer()

	// register our service
	cachely.RegisterCacheServer(s, &server{
		data: sync.Map{},
	})

	// Let the world know we are starting and where we are listening
	log.Printf("starting gRPC service on %s\n", lis.Addr())

	// start listening and responding
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
