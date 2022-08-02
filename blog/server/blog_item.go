package main

import (
	pb "github.com/shaikzhafir/udemy-go-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Content struct {
	BlogText string `bson:"blog_text"`
	BlogType string `bson:"blog_type"`
}

type BlogItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	AuthorId    string             `bson:"author_id""`
	Title       string             `bson:"title""`
	LastUpdated time.Time          `bson:"last_updated"`
	Content     Content            `bson:"content"`
}

func documentToBlog(data *BlogItem) *pb.Blog {
	return &pb.Blog{
		Id:          data.ID.Hex(),
		AuthorId:    data.AuthorId,
		Title:       data.Title,
		LastUpdated: timestamppb.New(data.LastUpdated),
		Content: &pb.Content{
			BlogText: data.Content.BlogText,
			BlogType: data.Content.BlogType,
		},
	}
}
