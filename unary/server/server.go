/*
   @Author: Bhaskar S
   @Blog:   https://www.polarsparc.com
   @Date:   28 Nov 2020
*/

package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"polarsparc.com/grpc/unary/greetpb"
	"time"
)

type server struct {}

func (s *server) Greet(_ context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("Received a Greet request with req: %v\n", req)

	name := req.GetName()
	message := getMessage(name)
	res := &greetpb.GreetResponse{
		Message: message,
	}

	return res, nil
}

const (
	addr = "127.0.0.1:20001"
)

func main() {
	log.Printf("Ready to start the Greet server on %s", addr)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to create listener on %s", addr)
	}

	srv := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(srv, &server{})

	if err = srv.Serve(lis); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getMessage(name string) string {
	hour := time.Now().Hour()

	msg := "Hello, " + name + ", "
	if hour < 12 {
		msg = msg + "Good Morning !!!"
	} else if hour < 16 {
		msg = msg + "Good Afternoon !!!"
	} else if hour < 21 {
		msg = msg + "Good Evening !!!"
	} else {
		msg = msg + "Good Night !!!"
	}

	return msg
}
