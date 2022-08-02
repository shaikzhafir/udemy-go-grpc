package main

import (
	"context"
	pb "github.com/shaikzhafir/udemy-go-grpc/blog/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"log"
	"time"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("inside create blog")

	blog := &pb.Blog{
		AuthorId:    "John",
		Title:       "What i ate today",
		LastUpdated: timestamppb.New(time.Now()),
		Content: &pb.Content{
			BlogText: "My first blog",
			BlogType: "happy feelings",
		},
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("error creating blog %v", err)
	}

	log.Printf("you have created blog with ID %s", res.Id)
	return res.Id
}

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("inside read blog")

	req := &pb.BlogId{Id: id}

	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Printf("Error happeend while reading %v", err)
	}

	log.Printf("blog was read %v\n", res)

	return res
}

func listBlog(c pb.BlogServiceClient) {
	log.Println("listblog invoked")

	stream, err := c.ListBlogs(context.Background(), &pb.Empty{})

	if err != nil {
		log.Fatalf("errir calling list blogs %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error reading from stream %v\n", err)
		}

		log.Println(res)
	}
}
