package iteration

import (
	"reflect"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"
	assertCorrectMessage(t, repeated, expected)
}

func TestSumSlice(t *testing.T) {
	sum := SumSlice([]int{1, 2, 3, 4, 5})
	expected := 15
	assertCorrectMessage(t, sum, expected)
}

func TestSumSlices(t *testing.T) {
	sum := SumSlices([]int{1, 2, 3}, []int{4, 4})
	expected := []int{6, 8}

	if !reflect.DeepEqual(sum, expected) {
		t.Errorf("got %v wanted %v", sum, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func assertCorrectMessage(t testing.TB, got any, want any) {
	t.Helper()
	if got != want {
		t.Errorf("expected '%v' but got '%v'", want, got)
	}
}
