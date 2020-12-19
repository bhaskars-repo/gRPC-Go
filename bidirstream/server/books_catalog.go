/*
   @Author: Bhaskar S
   @Blog:   https://www.polarsparc.com
   @Date:   19 Dec 2020
*/

package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"strings"
)

type Book struct {
	Title string
	ISBN string
}

type BooksCatalog map[string][]Book

type server struct {
	timeout int
	catalog BooksCatalog
}

func (s *server) Init() {
	l1 := []Book{{":The C++ Programming Language", "11111"},
		{"C++ Primer", "22222"},
		{"A Tour of C++", "33333"}}
	s.catalog["C++"] = l1
	l2 := []Book{{":The Go Programming Language", "44444"},
		{"Go in Practice", "55555"},
		{"Black Hat Go", "66666"}}
	s.catalog["GO"] = l2
	l3 := []Book{{":Effective Java", "77777"},
		{"Modern Java in Action", "88888"},
		{"Java: The Complete Reference", "99999"}}
	s.catalog["JAVA"] = l3
	l4 := []Book{{":Python Crash Course", "12121"},
		{"Learning Python", "34343"},
		{"Effective Python", "56565"}}
	s.catalog["PYTHON"] = l4
}

func (s *server) SetTimeout(t int) {
	s.timeout = t
}

func (s *server) GetBooks(topic string) ([]Book, error) {
	key := strings.ToUpper(topic)

	log.Printf("Books request for topic: %s\n", key)

	books := s.catalog[key]
	if books == nil {
		return nil, errors.New(fmt.Sprintf("No books for topic: %s", topic))
	}

	log.Printf("Books for key: %s = %v", key, books)

	return books, nil
}
