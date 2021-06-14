package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
	"sync"
)

// TODO: this function searches a slice to find a string value, to check for duplicates, inefficient, use a map to do this...
func sliceContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func traverse(url string, n *html.Node, resultSlice *[]string, mutex *sync.Mutex) {
	// this identifies anchor type html elements...
	if n.Type == html.ElementNode && n.Data == "a" {
		// ...and extracts their href values
		var urlStr = n.Attr[0].Val
		if strings.Contains(urlStr, "http") && strings.Contains(urlStr, url) {
			fmt.Println(urlStr)
			// we only want to write to the slice if it doesn't already contain this value, so we don't end up in an infinite loop
			if !sliceContains(*resultSlice, urlStr) {
				mutex.Lock()
				*resultSlice = append(*resultSlice, urlStr)
				mutex.Unlock()

				// then we recursively kick off the crawl again with the new page value
				parsePage(urlStr, resultSlice, mutex)
			}
		}
	}

	// this allows us to traverse all the child elements of the current page element, recursively
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverse(url, c, resultSlice, mutex)
	}
}

func parsePage(url string, resultSlice *[]string, mutex *sync.Mutex) {
	// this function parses each page's html to look for the anchor links
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			os.Exit(0)
		}
	}()

	doc, err := html.Parse(res.Body)
	traverse(url, doc, resultSlice, mutex)
}

func main() {
	args := os.Args[1:]
	var resultSlice []string
	var mutex = sync.Mutex{}
	parsePage(args[0], &resultSlice, &mutex)
	fmt.Printf("\n\n%+v\n", resultSlice)
}
