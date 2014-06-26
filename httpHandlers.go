package gotourexercises

import (
	"fmt"
	"net/http"
)

type String string

func (s String) ServeHTTP(
	w http.ResponseWriter,
	req *http.Request) {
	fmt.Fprint(w, s)
}

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s *Struct) ServeHTTP(
	w http.ResponseWriter,
	req *http.Request) {
	fmt.Fprint(w, s)
}