package gotourexercises

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStringHttpHandler(t *testing.T) {
	ts := httptest.NewServer(http.Handler(String("I'm a frayed knot.")))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("Something went wrong: %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if string(body) != "I'm a frayed knot." {
		t.Errorf("Wrong response  %s", body)
	}
}

func TestStructHttpHandler(t *testing.T) {
	ts := httptest.NewServer(http.Handler(&Struct{"Hello", ":", "Gophers!"}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("Something went wrong: %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if string(body) != "&{Hello : Gophers!}" {
		t.Errorf("Wrong response  %s", body)
	}
}
