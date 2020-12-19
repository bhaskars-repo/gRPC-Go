/*
   @Author: Bhaskar S
   @Blog:   https://www.polarsparc.com
   @Date:   19 Dec 2020
*/

package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
	"os"
	"polarsparc.com/grpc/bidirstream/bookspb"
	"strconv"
	"time"
)

func (s *server) RecommendedBooks(stream bookspb.BookRecommendService_RecommendedBooksServer) error {
	for {
		req, err := stream.Recv()
		if err == nil {
			books, err := s.GetBooks(req.Topic)
			if err == nil {
				for _, bk := range books {
					res := &bookspb.BookRecommendResponse{Topic: req.Topic, Title: bk.Title, Isbn: bk.ISBN}
					err = stream.Send(res)
					if err != nil {
						code := status.Code(err)
						if code == codes.Canceled {
							log.Println("ERROR - Deadline breached by BookRecommendService at localhost:20004")
						} else {
							log.Printf("Failed to send from BookRecommendService at localhost:20004: %v", err)
						}
						return err
					}
					time.Sleep(time.Duration(s.timeout) * time.Millisecond)
				}
			} else {
				log.Printf("Encountered an error on the server: %v", err)
				return status.Errorf(codes.InvalidArgument, err.Error())
			}
		} else if err == io.EOF {
			return nil
		} else {
			code := status.Code(err)
			if code == codes.Canceled {
				log.Println("ERROR - Deadline breached by BookRecommendService at localhost:20004")
			} else {
				log.Printf("Encountered an error for BookRecommendService at localhost:20004: %v\n", err)
			}
			return err
		}
	}
}

func main()  {
	tm := 100
	if len(os.Args) == 2 {
		tm, _ = strconv.Atoi(os.Args[1])
	}
	cs := &server{
		catalog: BooksCatalog{},
	}
	cs.Init()
	cs.SetTimeout(tm)

	log.Printf("Ready to start the BookRecommendService server with timeout %d...\n", cs.timeout)

	lis, err := net.Listen("tcp", "localhost:20004")
	if err != nil {
		log.Fatalf("Failed to create listener on localhost:20004")
	}

	srv := grpc.NewServer()

	bookspb.RegisterBookRecommendServiceServer(srv, cs)

	if err = srv.Serve(lis); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
