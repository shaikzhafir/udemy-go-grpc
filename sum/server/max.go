package main

import (
	pb "github.com/shaikzhafir/udemy-go-grpc/sum/proto"
	"io"
	"log"
)

func (s *Server) Max(stream pb.MaxService_MaxServer) error {
	log.Println("inside do Max stream")

	var maxNumber int64 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error reading client stream")
		}

		if req.StreamingNumber > maxNumber {
			maxNumber = req.StreamingNumber
		}

		err = stream.Send(&pb.MaxResponse{
			Result: float64(maxNumber),
		})

		if err != nil {
			log.Fatalf("error sending data to client %v\n", err)
		}

	}
}
