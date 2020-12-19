/*
   @Author: Bhaskar S
   @Blog:   https://www.polarsparc.com
   @Date:   12 Dec 2020
*/

package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
	"polarsparc.com/grpc/clientstream/quotepb"
)

func (s *server) GetBestQuote(stream quotepb.BestQuoteService_GetBestQuoteServer) error {
	var provider string
	var price float64

	for {
		req, err := stream.Recv()
		if err == nil {
			pq, err := s.GetProviderQuote(req.Provider, req.Age)
			if err == nil {
				if provider == "" || price > pq.Price {
					provider = pq.Provider
					price = pq.Price
				}
			} else {
				log.Printf("Encountered an error on the server: %v", err)
				return status.Errorf(codes.InvalidArgument, err.Error())
			}
		} else if err == io.EOF {
			// Received all client requests
			log.Printf("===> Best quote Provider: %s, Price: %.03f\n", provider, price)

			return stream.SendAndClose(&quotepb.BestQuoteResponse{
				Provider: provider,
				Price: price,
			})
		} else {
			log.Printf("Encountered an error for BestQuote at localhost:20003: %v\n", err)
			return status.Errorf(codes.FailedPrecondition, err.Error())
		}
	}
}

func main()  {
	qs := &server{
		cache: QuotesCache{},
	}
	qs.Init()

	log.Println("Ready to start the BestQuote server...")

	lis, err := net.Listen("tcp", "localhost:20003")
	if err != nil {
		log.Fatalf("Failed to create listener on localhost:20003")
	}

	srv := grpc.NewServer()

	quotepb.RegisterBestQuoteServiceServer(srv, qs)

	if err = srv.Serve(lis); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
