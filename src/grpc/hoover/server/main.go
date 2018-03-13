package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/golang/protobuf/ptypes/duration"
	"github.com/gopherguides/training/distributed-systems/grpc/src/hoover/hoover"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Port to communicate on
const port = ":50051"

// server is simply an empty struct for this example. In production, you would likely have fields as well.
type server struct{}

func (s *server) Get(ctx context.Context, in *hoover.GetRequest) (*hoover.GetReply, error) {
	// Check the context to make sure we haven't canceled or timed out.
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		return nil, status.New(codes.DeadlineExceeded, "context deadline exceeded").Err()
	default:
	}

	url := in.GetUrl()
	log.Printf("retrieving %s\n", url)

	// start a timer for this request
	begin := time.Now()

	// Retrieve the site
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	elapsed := durationToProtoDuration(time.Since(begin))

	defer resp.Body.Close()
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	log.Printf("finished retrieving %s\n", url)

	return &hoover.GetReply{
		ResponseCode: int32(resp.StatusCode),
		Body:         string(body),
		Elapsed:      elapsed,
	}, status.New(codes.OK, "").Err()
}

func main() {
	// Get a port to communicate on
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a new grpc server
	s := grpc.NewServer()

	// register our server
	hoover.RegisterServiceServer(s, &server{})

	// Let the world know we are starting and where we are listening
	log.Printf("starting gRPC service on %s\n", lis.Addr())

	// start serving
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// section: duration
func durationToProtoDuration(d time.Duration) *duration.Duration {
	seconds := int64(d) / int64(time.Second)
	nanos := int64(d) - (seconds * int64(time.Second))
	return &duration.Duration{Seconds: seconds, Nanos: int32(nanos)}
}

// section: duration
