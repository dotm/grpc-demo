package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	demoPB "dotm/grpc-demo/services/demo"

	"google.golang.org/grpc"
)

// flag.Parse() needs to be called in main function
var (
	port = flag.Int("port", 50051, "gRPC server port")
)

type demoServer struct {
	demoPB.UnimplementedDemoServer
}

func (s *demoServer) SimpleUnaryMethod(
	ctx context.Context, demoRequest *demoPB.DemoRequest,
) (*demoPB.DemoResponse, error) {
	responseMessage := fmt.Sprintf("Acknowledged: %s", demoRequest.GetRequestMessage())
	return &demoPB.DemoResponse{ResponseMessage: responseMessage}, nil
}

func (s *demoServer) ServerStreamingMethod(
	demoRequest *demoPB.DemoRequest, stream demoPB.Demo_ServerStreamingMethodServer,
) (err error) {
	for i := 0; i <= 3; i++ {
		responseMessage := fmt.Sprintf("Acknowledged %d: %s", i, demoRequest.GetRequestMessage())
		err = stream.Send(&demoPB.DemoResponse{ResponseMessage: responseMessage})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *demoServer) ClientStreamingMethod(
	stream demoPB.Demo_ClientStreamingMethodServer,
) (err error) {
	requestMessages := []string{}
	for {
		demoRequest, err := stream.Recv()
		if err == io.EOF {
			responseMessage := fmt.Sprintf("Acknowledged: %s", strings.Join(requestMessages, ", "))
			err = stream.SendAndClose(&demoPB.DemoResponse{ResponseMessage: responseMessage})
			return err
		}
		if err != nil {
			return err
		}

		requestMessages = append(requestMessages, demoRequest.RequestMessage)
	}
}

func (s *demoServer) BidirectionalStreamingMethod(
	stream demoPB.Demo_BidirectionalStreamingMethodServer,
) (err error) {
	i := 0
	for {
		demoRequest, err := stream.Recv()
		if err == io.EOF {
			return nil //close server's stream
		}
		if err != nil {
			return err
		}

		responseMessage := fmt.Sprintf("Acknowledged %d: %s", i, demoRequest.GetRequestMessage())
		err = stream.Send(&demoPB.DemoResponse{ResponseMessage: responseMessage})
		if err != nil {
			return err
		}
		i++
	}
}

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to create listener: %v", err)
	}

	s := grpc.NewServer()
	demoPB.RegisterDemoServer(s, &demoServer{})

	fmt.Printf("server listening at port %d\n", *port)
	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
