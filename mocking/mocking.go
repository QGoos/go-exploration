package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	countdownStartVal  = 3
	countdownFinalWord = "Go!"
	write              = "write"
	sleep              = "sleep"
)

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func Countdown(buff io.Writer, sleeper Sleeper) {
	for i := countdownStartVal; i > 0; i-- {
		fmt.Fprintln(buff, i)
		sleeper.Sleep()
	}
	fmt.Fprintln(buff, countdownFinalWord)
}

func CallCountdown() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
