package main

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
	"time"
)

const write = "write"
const sleep = "sleep"

type CountdownOperationsSpy struct {
	Calls []string
}

type SpyTime struct {
	durationSlept time.Duration
}

func TestCountdown(t *testing.T) {

	t.Run("Prints 3,2,1 Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &CountdownOperationsSpy{})

		got := buffer.String()
		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("Got '%q' and want '%q'.", got, want)
		}
	})

	t.Run("Sleep before every char", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)
		want := []string{sleep, write, sleep, write, sleep, write, sleep, write}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("Wanted calls '%v' got '%v'.", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}

	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("Should have slept for '%v' but slept for '%v'.", sleepTime, spyTime.durationSlept)
	}
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func ExampleCountdown() {
	buffer := &bytes.Buffer{}
	Countdown(buffer, &CountdownOperationsSpy{})
	result := buffer.String()
	fmt.Println(result)
	// Output: 3
	//2
	//1
	//Go!
}
