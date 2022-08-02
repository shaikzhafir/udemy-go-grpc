package main

import (
	"context"
	"fmt"
	"github.com/shaikzhafir/udemy-go-grpc/greet/proto"
	"io"
	"log"
)

// Greet client send single request, server send single response
func (s *Server) Greet(ctx context.Context, in *proto.GreetRequest) (*proto.GreetResponse, error) {
	log.Printf("greet function invoked with %v\n", in)
	return &proto.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}

// GreetManyTimes client send single request, server send stream
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

// ClientStreamGreet client streaming messages, server sends single response
func (s *Server) ClientStreamGreet(stream proto.GreetService_ClientStreamGreetServer) error {
	log.Printf("inside server side of client stream")

	res := "Here are your messages:"

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&proto.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("error reading client stream %v", err)
		}
		res += fmt.Sprintf(" %s", req.FirstName)

	}

}

func (s *Server) GreetEveryone(stream proto.GreetService_GreetEveryoneServer) error {
	log.Println("greet everyone invoked")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error reading client stream")
		}

		res := "Hello " + req.FirstName + "!"
		err = stream.Send(&proto.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("error sending data to client %v\n", err)
		}

	}
}
