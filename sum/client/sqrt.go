package main

import (
	"context"
	pb "github.com/shaikzhafir/udemy-go-grpc/sum/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("doSqrt invoked")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{
		Number: n,
	})

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("error message from server %s\n", e.Message())

			if e.Code() == codes.InvalidArgument {
				log.Println("negative number probably sent")
				return
			}
		} else {
			log.Fatalf("non gRPC error %v", err)
		}
	}

	log.Printf("sqrt : %f\n", res.Result)

}
