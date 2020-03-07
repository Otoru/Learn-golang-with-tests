package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("compares speeds of two servers", func(t *testing.T) {
		serverWithSlow := makeServerWithDelay(20 * time.Millisecond)
		defer serverWithSlow.Close()

		serverWithSpeed := makeServerWithDelay(0 * time.Millisecond)
		defer serverWithSpeed.Close()

		slow := serverWithSlow.URL
		fast := serverWithSpeed.URL

		result, err := Racer(slow, fast)

		if err != nil {
			t.Fatalf("Did not expect an error but got one '%v'.", err)
		}

		if result != fast {
			t.Errorf("Got '%q' ad want '%q'.", result, fast)
		}
	})

	t.Run("Returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeServerWithDelay(25 * time.Millisecond)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)
		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeServerWithDelay(delay time.Duration) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	})

	return httptest.NewServer(handler)
}
