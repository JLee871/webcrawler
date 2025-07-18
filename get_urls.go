package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing base url: %v", err)
	}

	reader := strings.NewReader(htmlBody)

	htmlNodes, err := html.Parse(reader)
	if err != nil {
		return nil, fmt.Errorf("error parsing nodes: %v", err)
	}

	var output []string

	for node := range htmlNodes.Descendants() {
		if node.Type == html.ElementNode && node.DataAtom == atom.A {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					url, err := url.Parse(attr.Val)
					if err != nil {
						continue
					}
					resolvedURL := baseURL.ResolveReference(url)
					output = append(output, resolvedURL.String())
				}
			}
		}
	}

	return output, nil
}
