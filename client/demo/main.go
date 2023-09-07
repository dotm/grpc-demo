package main

import (
	"context"
	"flag"
	"io"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	demoPB "dotm/grpc-demo/services/demo"
)

// flag.Parse() needs to be called in main function
var (
	serverAddress = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
)

func main() {
	flag.Parse()

	connection, err := grpc.Dial(*serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server %s: %v", *serverAddress, err)
	}
	defer connection.Close()

	clientStub := demoPB.NewDemoClient(connection)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("calling SimpleUnaryMethod")
	demoResponse, err := clientStub.SimpleUnaryMethod(ctx, &demoPB.DemoRequest{RequestMessage: "Hi"})
	if err != nil {
		log.Fatalf("error calling SimpleUnaryMethod: %v", err)
	}
	log.Println(demoResponse.ResponseMessage)
	log.Println("")

	log.Println("calling ServerStreamingMethod")
	serverStream, err := clientStub.ServerStreamingMethod(ctx, &demoPB.DemoRequest{RequestMessage: "Yo"})
	if err != nil {
		log.Fatalf("error calling ServerStreamingMethod: %v", err)
	}
	for {
		demoResponse, err = serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error receiving from ServerStreamingMethod: %v", err)
		}
		log.Println(demoResponse.ResponseMessage)
	}
	log.Println("")

	log.Println("calling ClientStreamingMethod")
	requestMessages := []string{"Hey", "Hello", "AYE", "YOU HEAR ME!!!"}
	clientStream, err := clientStub.ClientStreamingMethod(ctx)
	if err != nil {
		log.Fatalf("error calling ClientStreamingMethod: %v", err)
	}
	for _, msg := range requestMessages {
		err = clientStream.Send(&demoPB.DemoRequest{RequestMessage: msg})
		if err != nil {
			log.Fatalf("error sending to ClientStreamingMethod: %v", err)
		}
	}
	demoResponse, err = clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error CloseAndRecv of ClientStreamingMethod: %v", err)
	}
	log.Println(demoResponse.ResponseMessage)
	log.Println("")

	log.Println("calling BidirectionalStreamingMethod")
	bidirectionalStream, err := clientStub.BidirectionalStreamingMethod(ctx)
	if err != nil {
		log.Fatalf("error calling BidirectionalStreamingMethod: %v", err)
	}
	wg := sync.WaitGroup{}
	go func() {
		for {
			demoResponse, err = bidirectionalStream.Recv()
			if err == io.EOF {
				// read done
				break
			}
			if err != nil {
				log.Fatalf("error receiving from BidirectionalStreamingMethod: %v", err)
			}
			log.Printf("received: %s\n", demoResponse.ResponseMessage)
			wg.Done()
		}
		bidirectionalStream.CloseSend()
	}()
	for i := 0; i < len(requestMessages); i++ {
		log.Printf("sending: %s\n", requestMessages[i])
		err = bidirectionalStream.Send(&demoPB.DemoRequest{RequestMessage: requestMessages[i]})
		if err != nil {
			log.Fatalf("error sending to BidirectionalStreamingMethod: %v", err)
		}
		wg.Add(1)

		if i%2 == 1 {
			//send two messages and then wait for response
			wg.Wait()
		}
	}
	log.Println("")

	log.Println("terminating client")
}
