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
