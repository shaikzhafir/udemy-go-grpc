package main

import (
	pb "github.com/shaikzhafir/udemy-go-grpc/primes/proto"
	"log"
)

func (s *Server) Primes(in *pb.PrimesRequest, stream pb.PrimesService_PrimesServer) error {
	log.Printf("primes invoked with %v\n", in)

	var k int64 = 2
	var N = in.PrimeNumber
	for N > 1 {
		if N%k == 0 {
			err := stream.Send(&pb.PrimesResponse{
				Result: k,
			})
			if err != nil {
				return err
			}
			N = N / k // divide N by k so that we have the rest of the number left.
		} else {
			k = k + 1
		}
	}

	return nil
}
