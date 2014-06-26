package gotourexercises

import (
	"net/http"
	"testing"
)

func TestHttpHandler(t *testing.T) {
	// your http.Handle calls here
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	
	/*
	err := http.ListenAndServe("localhost:4001", nil)
	if err != nil {
		t.Errorf("Something went wrong: %v", err)
	}
	*/
}