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

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is jsut a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is jsut a test"

		assertCorrectMessage(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unkown")
		want := ErrNotFound

		assertError(t, err, want)
	})

}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")

		want := "this is just a test"
		got, err := dictionary.Search("test")

		assertError(t, err, nil)
		assertCorrectMessage(t, got, want)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		got, _ := dictionary.Search("test")

		assertError(t, err, ErrWordExists)
		assertCorrectMessage(t, got, definition)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("word exists", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		dictionary.Update(word, newDefinition)
		got, err := dictionary.Search("test")

		assertError(t, err, nil)
		assertCorrectMessage(t, got, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("word exists", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definition"}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		assertError(t, err, ErrNotFound)
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}

		AssertEqual(t, Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}

func TestBadBank(t *testing.T) {
	transactions := []Transaction{
		{
			From: "Chris",
			To:   "Riya",
			Sum:  100.0,
		},
		{
			From: "Adil",
			To:   "Chris",
			Sum:  25.0,
		},
	}

	AssertEqual(t, BalanceFor(transactions, "Riya"), 100.0)
	AssertEqual(t, BalanceFor(transactions, "Chris"), -75.0)
	AssertEqual(t, BalanceFor(transactions, "Adil"), -25.0)
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

func assertError(t testing.TB, err error, want error) {
	t.Helper()
	if err != want {
		t.Errorf("got error %q want %q", err, want)
	}
}

func AssertEqual(t *testing.T, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual(t *testing.T, got, want any) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %v", got)
	}
}
