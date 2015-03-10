package gotourexercises

import (
	"fmt"
	"net/http"
	"testing"
)

func PrepareHandlers() {
	// your http.Handle calls here
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", r.URL.Query()["code"])
	})
}

func TestHttpHandler(t *testing.T) {
	PrepareHandlers()

	err := http.ListenAndServe("localhost:4001", nil)
	if err != nil {
		t.Errorf("Something went wrong: %v", err)
	}

}
