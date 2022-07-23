package main

import (
	"github.com/shaikzhafir/udemy-go-grpc/sum/proto"
	"github.com/shaikzhafir/udemy-go-grpc/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial(utils.GetAddr(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()
	// add logic to consume request

	c := proto.NewSumServiceClient(conn)

	doSum(c)
}
