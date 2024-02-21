package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("print", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Charles")

		got := buffer.String()
		want := "Hello, Charles"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected '%v' but got '%v'", want, got)
	}
}
