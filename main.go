package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawURL := args[0]

	const defaultMaxConcurrency = 3
	const defaultMaxPages = 100

	maxConcurrency := defaultMaxConcurrency
	maxPages := defaultMaxPages
	if len(args) > 1 {
		num, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("bad arg. usage: <base_url> [max_concurrency {3}] [max_pages {100}]")
		}
		maxConcurrency = num
	}
	if len(args) > 2 {
		num, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("bad arg. usage: <base_url> [max_concurrency {3}] [max_pages {100}]")
		}
		maxPages = num
	}

	cfg, err := configure(rawURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("error configuring: %v", err)
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %v\n", rawURL)
	cfg.wg.Add(1)
	go cfg.crawlPage(rawURL)
	cfg.wg.Wait()

	for key, value := range cfg.pages {
		fmt.Printf("%v: %v\n", key, value)
	}
}
