package main

import (
	"context"
	pb "github.com/shaikzhafir/udemy-go-grpc/sum/proto"
	"log"
)

func doSum(c pb.SumServiceClient) {
	log.Println("doGreet invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstInt:  34234,
		SecondInt: 43,
	})

	if err != nil {
		log.Fatalf("error consuming API %v\n", err)
	}

	log.Printf("Result is: %v", res.Result)
}
