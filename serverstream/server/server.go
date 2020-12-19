/*
   @Author: Bhaskar S
   @Blog:   https://www.polarsparc.com
   @Date:   04 Dec 2020
*/

package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"polarsparc.com/grpc/serverstream/currencypb"
	"time"
)

func (s *server) GetCurrencyRate(req *currencypb.CurrencyRateRequest, stream currencypb.CurrencyService_GetCurrencyRateServer) error {
	log.Printf("Received a CurrencyRate request with req: %v\n", req)

	from := req.From
	to := req.To

	rates, err := s.GetAgentRates(from, to)
	if err == nil {
		log.Printf("Rates from agents: %v\n", rates)
		for _, r := range rates {
			res := &currencypb.CurrencyRateResponse{Agent: r.Agent, From: from, To: to, Rate: r.Rate}
			stream.Send(res)
			time.Sleep(250 * time.Millisecond)
		}
		return nil
	}

	return err
}

func main()  {
	cs := &server{
		cache: RatesCache{},
	}
	cs.Init()

	log.Println("Ready to start the CurrencyRate server...")

	lis, err := net.Listen("tcp", "localhost:20002")
	if err != nil {
		log.Fatalf("Failed to create listener on localhost:20002")
	}

	srv := grpc.NewServer()

	currencypb.RegisterCurrencyServiceServer(srv, cs)

	if err = srv.Serve(lis); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
