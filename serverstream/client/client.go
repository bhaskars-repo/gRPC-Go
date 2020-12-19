/*
   @Author: Bhaskar S
   @Blog:   https://www.polarsparc.com
   @Date:   04 Dec 2020
*/

package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"polarsparc.com/grpc/serverstream/currencypb"
)

func main() {
	log.Println("Ready to start the CurrencyRate client...")

	conn, err := grpc.Dial("localhost:20002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to localhost:20002")
	}
	defer conn.Close()

	cl := currencypb.NewCurrencyServiceClient(conn)

	// Success
	req := &currencypb.CurrencyRateRequest{From: "usd",
		To: "eur"}
	stream, err := cl.GetCurrencyRate(context.Background(), req)
	if err != nil {
		log.Fatalf("[1] Failed to send CurrencyRate request to localhost:20002")
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("[1] Received an error from CurrencyRate at localhost:20002: %v", err)
		}
		log.Printf("[1] ===> Agent: %s, Rate: %.03f\n", res.Agent, res.Rate)
	}

	// Error
	req2 := &currencypb.CurrencyRateRequest{From: "usd",
		To: "jpy"}
	stream2, err := cl.GetCurrencyRate(context.Background(), req2)
	if err != nil {
		log.Fatalf("[2] Failed to send CurrencyRate request to localhost:20002")
	}
	for {
		res, err := stream2.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("[2] Received and error from CurrencyRate at localhost:20002: %v", err)
		}
		log.Printf("[2] ===> Agent: %s, Rate: %.03f\n", res.Agent, res.Rate)
	}
}
