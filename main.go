package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawURL := args[0]

	const maxConcurrency = 3
	cfg, err := configure(rawURL, maxConcurrency)
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
