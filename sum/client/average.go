package main

import (
	"context"
	pb "github.com/shaikzhafir/udemy-go-grpc/sum/proto"
	"log"
)

func doAverageClientStreaming(c pb.AverageServiceClient) {
	log.Println("average client streaming invoked")

	reqs := []*pb.AverageRequest{
		{StreamingNumber: 4},
		{StreamingNumber: 2},
		{StreamingNumber: 7},
		{StreamingNumber: 54},
	}

	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalf("error starting client stream")
	}

	for _, req := range reqs {
		log.Printf("sending req %v", req)
		err := stream.Send(req)

		if err != nil {
			log.Fatalf("error sending stream %v", err)
		}

	}
	//stream done, close and receive response from server
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error receiving response from server: %v ", err)
	}

	log.Printf("Client stream ended with response: %v", res.Result)

}
