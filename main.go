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
	fmt.Printf("starting crawl of: %v\n", rawURL)

	pages := make(map[string]int)
	crawlPage(rawURL, rawURL, pages)

	for key, value := range pages {
		fmt.Printf("%v: %v\n", key, value)
	}
}
