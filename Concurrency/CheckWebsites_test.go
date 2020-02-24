package concurrency

import (
	"reflect"
	"testing"
	"time"
	"fmt"
)

func mockWebsiteChecker(url string) bool {
	if url == "wtf://randon.site" {
		return false
	}
	return true
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://facebook.com",
		"http://google.com",
		"wtf://randon.site",
	}

	want := map[string]bool{
		"http://google.com": true,
		"http://facebook.com": true,
		"wtf://randon.site": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted '%v' and got '%v'.", want, got)
	}
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

func ExampleCheckWebsites() {
	urls := []string{
		"http://facebook.com",
		"http://google.com",
		"wtf://randon.site",
	}

	result := CheckWebsites(mockWebsiteChecker, urls)

	fmt.Println(result)
	// Output: map[http://facebook.com:true http://google.com:true wtf://randon.site:false]
}
