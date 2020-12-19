/*
   @Author: Bhaskar S
   @Blog:   https://www.polarsparc.com
   @Date:   04 Dec 2020
*/

package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"strings"
)

type CurrencyRate struct {
	Agent string
	Rate float64
}

type RatesCache map[string][]CurrencyRate

type server struct {
	cache RatesCache
}

func (s *server) Init() {
	l1 := []CurrencyRate{{Agent: "Alice", Rate: 1.30}, {Agent: "Bob", Rate: 1.302},{Agent: "Dave", Rate: 1.31}}
	s.cache["USD:CAD"] = l1
	l2 := []CurrencyRate{{Agent: "Alice", Rate: 0.85}, {Agent: "Charlie", Rate: 0.84}}
	s.cache["USD:EUR"] = l2
	l3 := []CurrencyRate{{Agent: "Bob", Rate: 0.75}, {Agent: "Charlie", Rate: 0.751},{Agent: "Eve", Rate: 0.74}}
	s.cache["USD:GBP"] = l3
}

func (s *server) GetAgentRates(from string, to string) ([]CurrencyRate, error) {
	key := strings.ToUpper(from + ":" + to)

	log.Printf("Currency rate request for key: %s\n", key)

	rates := s.cache[key]
	if rates == nil {
		return nil, errors.New(fmt.Sprintf("No rate for currency from: %s, to: %s", from, to))
	}

	log.Printf("Currency rates for key: %s = %v", key, rates)

	return rates, nil
}
