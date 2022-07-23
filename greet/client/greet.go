package main

import (
	"context"
	pb "github.com/shaikzhafir/udemy-go-grpc/greet/proto"
	"io"
	"log"
	"time"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "John",
	})

	if err != nil {
		log.Fatalf("error consuming API %v\n", err)
	}

	log.Printf("Greetings: %v", res.Result)
}

// send single message, receive a stream as response
func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("do greet many times invoked")

	stream, err := c.GreetManyTimes(context.Background(), &pb.GreetRequest{
		FirstName: "John",
	})

	if err != nil {
		log.Fatalf("error while calling greet many times %v", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error in message stream %v", err)
		}

		log.Printf("Greet many times %v", msg)

	}

}

//send a stream of messages, receive single response
func doClientStreamingGreet(c pb.GreetServiceClient) {
	log.Println("client streaming greet invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "This"},
		{FirstName: "Is"},
		{FirstName: "Poop"},
	}

	stream, err := c.ClientStreamGreet(context.Background())

	if err != nil {
		log.Fatalf("error starting client stream: %v ", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req %v", req)
		err := stream.Send(req)

		if err != nil {
			log.Fatalf("error sending stream %v", err)
		}

		// pause 1 second before sending subsequent message in stream
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error receiving response from server: %v ", err)
	}

	log.Printf("Client stream ended with response: %v", res.Result)

}
