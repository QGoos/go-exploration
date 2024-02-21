package mocking

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("countdown", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := "3\n2\n1\nGo!\n"
		expected_count := 3

		assertCorrectMessage(t, got, want)
		assertCorrectSleeperCount(t, *spySleeper, expected_count)
	})
	t.Run("alternating events", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write, sleep,
			write, sleep,
			write, sleep,
			write,
		}

		assertDeepCorrectMessage(t, spySleepPrinter, want)
	})
}

func assertCorrectMessage(t testing.TB, got any, want any) {
	t.Helper()
	if got != want {
		t.Errorf("expected '%v' but got '%v'", want, got)
	}
}

func assertCorrectSleeperCount(t testing.TB, sleeper SpySleeper, count int) {
	t.Helper()
	if sleeper.Calls != count {
		t.Errorf("not enough calls to sleeper, want %d but got %d", count, sleeper.Calls)
	}
}

func assertDeepCorrectMessage(t testing.TB, ssp *SpyCountdownOperations, want []string) {
	t.Helper()
	if !reflect.DeepEqual(want, ssp.Calls) {
		t.Errorf("expected '%v' but got '%v'", want, ssp)
	}
}
