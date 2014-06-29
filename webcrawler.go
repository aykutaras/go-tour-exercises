package gotourexercises

import (
    "fmt"
)

type Fetcher interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    Fetch(url string) (body string, urls []string, err error)
}

type sharedCrawl struct {
	Fetcher
	mapAccess chan map[string]bool
	printAccess chan bool
}

func (c *sharedCrawl) c2(url string, depth int, pageDone chan bool) {
	if depth <= 0 {
		pageDone <- true
        return
    }
    
    body, urls, err := c.Fetch(url)
    if err != nil {
		<-c.printAccess
        fmt.Println(err)
		c.printAccess<-true
		
		pageDone<-true
        return
    }
	
	<-c.printAccess
    fmt.Printf("found: %s %q\n", url, body)
	c.printAccess<-true
	
	uDone := make(chan bool)
	uCount := 0
	
	m := <-c.mapAccess
    for _, u := range urls {
        if !m[u] {
			m[u] = true
			uCount++
			go c.c2(u, depth-1, uDone)
        }
    }
	c.mapAccess<-m
	
	for ; uCount >0; uCount-- {
		<-uDone
	}
	
	pageDone<-true
    return
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher)bool {
	c := &sharedCrawl {
		fetcher,
		make(chan map[string]bool, 1),
		make(chan bool, 1),
	}
	
	c.mapAccess <- map[string]bool {url:true}
	c.printAccess <- true
	done := make(chan bool)
	go c.c2(url, depth, done)
	<-done
	
	return true
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
    body string
    urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
    if res, ok := f[url]; ok {
        return res.body, res.urls, nil
    }
    return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
    "http://golang.org/": &fakeResult{
        "The Go Programming Language",
        []string{
            "http://golang.org/pkg/",
            "http://golang.org/cmd/",
        },
    },
    "http://golang.org/pkg/": &fakeResult{
        "Packages",
        []string{
            "http://golang.org/",
            "http://golang.org/cmd/",
            "http://golang.org/pkg/fmt/",
            "http://golang.org/pkg/os/",
        },
    },
    "http://golang.org/pkg/fmt/": &fakeResult{
        "Package fmt",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
    "http://golang.org/pkg/os/": &fakeResult{
        "Package os",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
}