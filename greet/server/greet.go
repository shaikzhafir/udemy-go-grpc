package main

import (
	"context"
	"fmt"
	"github.com/shaikzhafir/udemy-go-grpc/greet/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, in *proto.GreetRequest) (*proto.GreetResponse, error) {
	log.Printf("greet function invoked with %v\n", in)
	return &proto.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}

func (s *Server) GreetManyTimes(in *proto.GreetRequest, stream proto.GreetService_GreetManyTimesServer) error {
	log.Printf("greet many times function invoked with %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, Response from server %d", in.FirstName, i)
		err := stream.Send(&proto.GreetResponse{
			Result: res,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
