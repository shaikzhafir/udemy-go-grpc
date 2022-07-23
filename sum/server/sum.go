package main

import (
	"context"
	"github.com/shaikzhafir/udemy-go-grpc/sum/proto"
	"log"
)

func (s *Server) Sum(ctx context.Context, in *proto.SumRequest) (*proto.SumResponse, error) {
	log.Printf("request is %v", in)

	return &proto.SumResponse{
		Result: in.FirstInt + in.SecondInt}, nil
}
