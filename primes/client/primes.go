package main

import (
	"context"
	pb "github.com/shaikzhafir/udemy-go-grpc/primes/proto"
	"io"
	"log"
)

func streamPrimes(c pb.PrimesServiceClient) {
	log.Println("stream primes invoked")

	stream, err := c.Primes(context.Background(), &pb.PrimesRequest{
		PrimeNumber: 120,
	})

	if err != nil {
		log.Fatalf("error while calling primes %v", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while streaming primes %v", err)
		}

		log.Printf("%d\n", msg.Result)
	}
}
