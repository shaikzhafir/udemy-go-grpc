package main

import (
	"context"
	pb "github.com/shaikzhafir/udemy-go-grpc/sum/proto"
	"io"
	"log"
	"time"
)

func doMaxClientStream(c pb.MaxServiceClient) {
	log.Println("doMaxClientStream invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("error starting client stream")
	}

	reqs := []*pb.MaxRequest{
		{StreamingNumber: 1},
		{StreamingNumber: 5},
		{StreamingNumber: 4},
		{StreamingNumber: 9},
	}

	waitc := make(chan struct{})

	// go func for client streaming messages and sending
	go func() {
		for _, req := range reqs {
			log.Println("sending request %s", req.StreamingNumber)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			msg, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("error in message stream %v", err)
				break
			}

			log.Printf("Max number so far %v", msg)
		}
		close(waitc)
	}()

	<-waitc

}
