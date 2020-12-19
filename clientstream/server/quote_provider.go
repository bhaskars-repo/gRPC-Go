/*
   @Author: Bhaskar S
   @Blog:   https://www.polarsparc.com
   @Date:   12 Dec 2020
*/

package main

import (
	"errors"
	"fmt"
	"log"
)

type ProviderQuote struct {
	Provider string
	AgeLow int32
	AgeHigh int32
	Price float64
}

func (pq ProviderQuote) inRange(age int32) bool {
	if age >= pq.AgeLow && age <= pq.AgeHigh {
		return true
	}
	return false
}

type QuotesCache map[string][]ProviderQuote

type server struct {
	cache QuotesCache
}

func (s *server) Init() {
	l1 := []ProviderQuote{{Provider: "Alice", AgeLow: 20, AgeHigh: 30, Price: 1000.0},
        {Provider: "Alice", AgeLow: 31, AgeHigh: 45, Price: 1500.0},
		{Provider: "Alice", AgeLow: 46, AgeHigh: 55, Price: 2000.0},
	}
	s.cache["Alice"] = l1

	l2 := []ProviderQuote{{Provider: "Bob", AgeLow: 20, AgeHigh: 30, Price: 1100.0},
		{Provider: "Bob", AgeLow: 31, AgeHigh: 45, Price: 1475.0},
		{Provider: "Bob", AgeLow: 46, AgeHigh: 55, Price: 1950.0},
	}
	s.cache["Bob"] = l2

	l3 := []ProviderQuote{{Provider: "Charlie", AgeLow: 20, AgeHigh: 30, Price: 975.0},
		{Provider: "Charlie", AgeLow: 31, AgeHigh: 45, Price: 1525.0},
		{Provider: "Charlie", AgeLow: 46, AgeHigh: 55, Price: 2050.0},
	}
	s.cache["Charlie"] = l3
}

func (s *server) GetProviderQuote(provider string, age int32) (*ProviderQuote, error) {
	log.Printf("Request for provider: %s, age: %d", provider, age)

	var pq *ProviderQuote

	quotes := s.cache[provider]
	if quotes == nil {
		return nil, errors.New(fmt.Sprintf("Specified provider %s invalid", provider))
	}

	for _, e := range quotes {
		if e.inRange(age) {
			pq = &e
			break
		}
	}

	if pq == nil {
		return nil, errors.New(fmt.Sprintf("No Quote for the specified provider %s and age %d", provider, age))
	}

	log.Printf("Provider quote for %s and age %d - %.02f\n", provider, age, pq.Price)

	return pq, nil
}
