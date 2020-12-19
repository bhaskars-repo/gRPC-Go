/*
   @Author: Bhaskar S
   @Blog:   https://www.polarsparc.com
   @Date:   28 Nov 2020
*/

package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"polarsparc.com/grpc/unary/greetpb"
)

const (
	addr = "127.0.0.1:20001"
)

func main()  {
	log.Println("Ready to start the Greet client...")

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to %s", addr)
	}
	defer conn.Close()

	cl := greetpb.NewGreetServiceClient(conn)

	req := &greetpb.GreetRequest{
		Name: "Alice",
	}
	res, err := cl.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to send Greet request to %s [%v]", addr, err)
	}

	log.Printf("%s\n", res.Message)
}
