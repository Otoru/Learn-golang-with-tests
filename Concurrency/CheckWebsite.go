package concurrency

import "net/http"

// CheckWebsite returns true if the URL returns a '200 Ok' status code
func CheckWebsite(url string) bool {
	response, err := http.Head(url)
	if err != nil {
		return false
	}

	if response.StatusCode == http.StatusOK {
		return true
	}

	return false
}