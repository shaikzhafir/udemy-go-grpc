package main

import (
	"github.com/shaikzhafir/udemy-go-grpc/sum/proto"
	"io"
	"log"
)

func (s *Server) Average(stream proto.AverageService_AverageServer) error {
	log.Println("isnide server side of client stream")

	var totalSum int64 = 0
	var count int64 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			average := float64(totalSum) / float64(count)
			return stream.SendAndClose(&proto.AverageResponse{
				Result: average,
			})
		}

		if err != nil {
			log.Fatalf("error reading client stream")
		}
		count += 1
		totalSum += req.StreamingNumber

	}

}
