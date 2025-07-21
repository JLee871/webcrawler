package main

import (
	"fmt"
	"net/url"
	"time"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	current, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("error parsing current url: %v\n", err)
		return
	}

	base, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("error parsing base url: %v\n", err)
		return
	}

	if base.Hostname() != current.Hostname() {
		return
	}

	normalized, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normalizing current url: %v\n", err)
		return
	}

	val, ok := pages[normalized]
	if ok {
		val += 1
		return
	} else {
		pages[normalized] = 1
		fmt.Println(normalized)
	}

	time.Sleep(5 * time.Second)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error getting html: %v\n", err)
		return
	}

	urls, err := getURLsFromHTML(html, rawBaseURL)
	if err != nil {
		fmt.Printf("error getting urls from html: %v\n", err)
		return
	}

	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}
}
