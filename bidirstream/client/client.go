/*
   @Author: Bhaskar S
   @Blog:   https://www.polarsparc.com
   @Date:   19 Dec 2020
*/

package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"polarsparc.com/grpc/bidirstream/bookspb"
	"time"
)

func main() {
	log.Println("Ready to start the BookRecommend client...")

	conn, err := grpc.Dial("localhost:20004", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to localhost:20004")
	}
	defer conn.Close()

	cl := bookspb.NewBookRecommendServiceClient(conn)

	wc := make(chan bool)

	// Test Case-1

	stream, err := cl.RecommendedBooks(context.Background())
	if err != nil {
		log.Fatalf("[1] Failed to create client stub to localhost:20004: %v", err)
	}

	go func() {
		err = stream.Send(&bookspb.BookRecommendRequest{
			Topic: "go",
		})
		if err != nil {
			log.Fatalf("[1] Failed to send request to localhost:20004: %v", err)
		}

		err = stream.Send(&bookspb.BookRecommendRequest{
			Topic: "Python",
		})
		if err != nil {
			log.Fatalf("[1] Failed to send request to localhost:20004: %v", err)
		}

		err = stream.CloseSend()
		if err != nil {
			log.Fatalf("[1] Failed to close send stream: %v", err)
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == nil {
				log.Printf("[1] Recommended book => topic: %s, title: %s, isbn: %s\n",
					res.Topic, res.Title, res.Isbn)
			} else if err == io.EOF {
				wc <-true
				break
			} else {
				log.Fatalf("[1] Encountered an error on receive from server: %v", err)
			}
		}
	}()

	// Test Case-1 Wait
	<-wc

	// Test Case-2 - Error case

	stream, err = cl.RecommendedBooks(context.Background())
	if err != nil {
		log.Fatalf("[2] Failed to create client stub to localhost:20004: %v", err)
	}

	go func() {
		err = stream.Send(&bookspb.BookRecommendRequest{
			Topic: "java",
		})
		if err != nil {
			log.Fatalf("[2] Failed to send request to localhost:20004: %v", err)
		}

		err = stream.Send(&bookspb.BookRecommendRequest{
			Topic: "rust",
		})
		if err != nil {
			log.Fatalf("[2] Failed to send request to localhost:20004: %v", err)
		}

		err = stream.CloseSend()
		if err != nil {
			log.Fatalf("[2] Failed to close send stream: %v", err)
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == nil {
				log.Printf("[2] Recommended book => topic: %s, title: %s, isbn: %s\n",
					res.Topic, res.Title, res.Isbn)
			} else if err == io.EOF {
				wc <-true
				break
			} else {
				st, ok := status.FromError(err)
				if ok {
					log.Printf("[2] Error - %s\n", st.Message())
				} else {
					log.Fatalf("[2] Unexpected failure from localhost:20004: %v", err)
				}
				wc <-true
				break
			}
		}
	}()

	// Test Case-2 Wait
	<-wc

	// Test Case-3 - Deadline/Timeout case

	duration := time.Now().Add(550 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), duration)
	defer cancel()

	stream, err = cl.RecommendedBooks(ctx)
	if err != nil {
		log.Fatalf("[3] Failed to create client stub to localhost:20004: %v", err)
	}

	go func() {
		err = stream.Send(&bookspb.BookRecommendRequest{
			Topic: "java",
		})
		if err != nil {
			log.Fatalf("[3] Failed to send request to localhost:20004: %v", err)
		}

		err = stream.Send(&bookspb.BookRecommendRequest{
			Topic: "python",
		})
		if err != nil {
			log.Fatalf("[3] Failed to send request to localhost:20004: %v", err)
		}

		err = stream.CloseSend()
		if err != nil {
			log.Fatalf("[3] Failed to close send stream: %v", err)
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == nil {
				log.Printf("[3] Recommended book => topic: %s, title: %s, isbn: %s\n",
					res.Topic, res.Title, res.Isbn)
			} else if err == io.EOF {
				wc <-true
				break
			} else {
				st, ok := status.FromError(err)
				if ok {
					log.Printf("[3] Error - %s\n", st.Message())
				} else {
					log.Fatalf("[3] Unexpected failure from localhost:20004: %v", err)
				}
				wc <-true
				break
			}
		}
	}()

	// Test Case-3 Wait
	<-wc
}
