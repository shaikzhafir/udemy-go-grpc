package main

import (
	pb "github.com/shaikzhafir/udemy-go-grpc/sum/proto"
	"github.com/shaikzhafir/udemy-go-grpc/utils"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pb.SumServiceServer
}

func main() {
	lis, err := net.Listen("tcp", utils.GetAddr())
	if err != nil {
		log.Fatalf("failed to listen on port %s", utils.GetAddr())
	}

	log.Printf("listning on %v", utils.GetAddr())

	s := grpc.NewServer()
	pb.RegisterSumServiceServer(s, &Server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve %v", err)
	}

}
