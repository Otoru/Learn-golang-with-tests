package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondsTimeout = 10 * time.Second

// Racer compares the response times of a and b, returning the fastest one with timeout after 10s
func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondsTimeout	)
}

// ConfigurableRacer compares the times response of a and b, returning the fastest one
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-get(a):
		return a, nil
	case <-get(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("Timed out waiting for '%s' and '%s'.", a, b)
	}
}

func get(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}