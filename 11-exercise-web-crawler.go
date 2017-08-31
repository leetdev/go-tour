// https://tour.golang.org/concurrency/10

package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Cache URLs already fetched
type UrlCache struct {
	urls map[string]bool
	mux  sync.Mutex
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache UrlCache, results chan string) {
	// Each crawl thread gets its own results channel,
	// which is closed automatically upon return
	defer close(results)

	// Start working on URL cache
	cache.mux.Lock()

	// Don't fetch the same URL twice.
	// Also respect the depth limit
	if cache.urls[url] || depth <= 0 {
		cache.mux.Unlock()
		return
	} else {
		cache.urls[url] = true
	}

	// Done working on URL cache
	cache.mux.Unlock()

	// Fetch
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		results <- fmt.Sprintf("%s", err)
		return
	}

	// Send result to channel
	results <- fmt.Sprintf("found: %s %q\n", url, body)

	// Create channels for additional results and then crawl the URLs
	result := make([]chan string, len(urls))
	for i, u := range urls {
		result[i] = make(chan string)
		go Crawl(u, depth-1, fetcher, cache, result[i])
	}

	// Print out the additional results to original results channel
	for i := range result {
		for s := range result[i] {
			results <- s
		}
	}
}

func main() {
	// Create results channel and cache map
	results := make(chan string)
	cache := UrlCache{urls: make(map[string]bool)}

	// Crawl
	go Crawl("http://golang.org/", 4, fetcher, cache, results)

	// Print results as they are returned from Go threads
	for i := range results {
		fmt.Print(i)
	}
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
