package gotourexercises

import "testing"

func TestWebCrawler(t *testing.T) {
    result := Crawl("http://golang.org/", 4, fetcher)
	
	if !result {
		t.Errorf("Test worked as expected")
	}
}