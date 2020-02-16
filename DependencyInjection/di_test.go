package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Vitor")

	got := buffer.String()
	want := "Hello Vitor"

	if got != want {
		t.Errorf("Got '%q' and want '%q'.", got, want)
	}
}

func ExampleGreet() {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Vitor")
	result := buffer.String()
	fmt.Println(result)
	// Output: Hello Vitor
}
