package main

import (
	"context"
	pb "github.com/shaikzhafir/udemy-go-grpc/blog/proto"
	"github.com/shaikzhafir/udemy-go-grpc/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pb.BlogServiceServer
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongouser:mongopassword@127.0.0.1:32017/"))

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("coneected to mongodb %v", client)
	err = client.Connect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", utils.GetAddr())
	if err != nil {
		log.Fatalf("failed to listen on port %s", utils.GetAddr())
	}

	log.Printf("listening on %s", utils.GetAddr())

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve %v", err)
	}

}
