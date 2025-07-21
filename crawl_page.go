package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	current, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("error parsing current url: %v\n", err)
		return
	}

	if cfg.baseURL.Hostname() != current.Hostname() {
		return
	}

	normalized, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normalizing current url: %v\n", err)
		return
	}

	isFirst := cfg.addPageVisit(normalized)
	if !isFirst {
		return
	}

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error getting html: %v\n", err)
		return
	}

	urls, err := getURLsFromHTML(html, cfg.baseURL)
	if err != nil {
		fmt.Printf("error getting urls from html: %v\n", err)
		return
	}

	for _, url := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}
