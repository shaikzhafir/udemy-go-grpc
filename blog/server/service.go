package main

import (
	"context"
	"fmt"
	pb "github.com/shaikzhafir/udemy-go-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("inside createBlog server: %v\n", in)

	data := BlogItem{
		AuthorId:    in.AuthorId,
		Title:       in.Title,
		LastUpdated: in.LastUpdated.AsTime(),
		Content: Content{
			BlogText: in.Content.BlogText,
			BlogType: in.Content.BlogType,
		},
	}

	res, err := collection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error: %v\n", err))
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"cannot convert to OID")
	}

	return &pb.BlogId{Id: oid.Hex()}, nil
}

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Printf("isnide ReadBlog server with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)

	// if id is an objectID, call from db
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID as objectID")
	}

	data := &BlogItem{}

	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)

	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"cannot find blog with ID provided")
	}

	return documentToBlog(data), nil

}

func (s *Server) ListBlogs(in *pb.Empty, stream pb.BlogService_ListBlogsServer) error {
	log.Printf("inside List blog")

	list, err := collection.Find(context.Background(), primitive.D{{}})

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error %v\n", err))
	}

	defer list.Close(context.Background())

	for list.Next(context.Background()) {
		data := &BlogItem{}
		err := list.Decode(data)

		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("error decoding data from mongodb %v\n", err))
		}

		stream.Send(documentToBlog(data))

	}

	if err = list.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("unknown internal error %v", err))
	}

	return nil

}
