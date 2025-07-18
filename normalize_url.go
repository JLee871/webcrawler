package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("error parsing url occured: %v", err)
	}

	path := u.Host + u.Path
	path = strings.ToLower(path)
	path = strings.TrimSuffix(path, "/")

	return path, nil
}
