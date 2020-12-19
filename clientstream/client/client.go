/*
   @Author: Bhaskar S
   @Blog:   https://www.polarsparc.com
   @Date:   12 Dec 2020
*/

package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
	"polarsparc.com/grpc/clientstream/quotepb"
)

func main() {
	log.Println("Ready to start the BestQuote client...")

	conn, err := grpc.Dial("localhost:20003", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to localhost:20003")
	}
	defer conn.Close()

	cl := quotepb.NewBestQuoteServiceClient(conn)

	// Case - 1

	stream, err := cl.GetBestQuote(context.Background())
	if err != nil {
		log.Fatalf("[1] Failed to create client stub to localhost:20003: %v", err)
	}

	err = stream.Send(&quotepb.BestQuoteRequest{
		Provider: "Bob",
		Age: 37,
	})
	if err != nil {
		log.Fatalf("[1] Failed to send request to localhost:20003: %v", err)
	}

	err = stream.Send(&quotepb.BestQuoteRequest{
		Provider: "Charlie",
		Age: 37,
	})
	if err != nil {
		log.Fatalf("[1] Failed to send request to localhost:20003: %v", err)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("[1] Received and error from BestQuote at localhost:20003: %v", err)
	}

	log.Printf("[1] Best quote from provider %s with price %.02f\n", res.Provider, res.Price)

	// Case - 2 - Error case

	stream, err = cl.GetBestQuote(context.Background())
	if err != nil {
		log.Fatalf("[2] Failed to create client stub to localhost:20003: %v", err)
	}

	err = stream.Send(&quotepb.BestQuoteRequest{
		Provider: "Alice",
		Age: 48,
	})
	if err != nil {
		log.Fatalf("[2-1] Failed to send request to localhost:20003: %v", err)
	}

	err = stream.Send(&quotepb.BestQuoteRequest{
		Provider: "Dave",
		Age: 48,
	})
	if err != nil {
		log.Fatalf("[2-2] Failed to send request to localhost:20003: %v", err)
	}

	res, err = stream.CloseAndRecv()
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			log.Printf("[2] Error - %s\n", st.Message())
		} else {
			log.Fatalf("[2] Unexpected failure from BestQuote at localhost:20003: %v", err)
		}
	}
}
