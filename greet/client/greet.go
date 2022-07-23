package main

import (
	"context"
	pb "github.com/shaikzhafir/udemy-go-grpc/greet/proto"
	"io"
	"log"
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
