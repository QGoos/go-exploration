package generics

import "testing"

func TestGenerics(t *testing.T) {
	t.Run("int stack", func(t *testing.T) {
		myStackOInts := new(Stack[int])

		//check stack empty
		AssertTrue(t, myStackOInts.IsEmpty())

		// add thing, check not empty
		myStackOInts.Push(1)
		AssertFalse(t, myStackOInts.IsEmpty())

		// add thing pop back again
		myStackOInts.Push(3)
		value, _ := myStackOInts.Pop()
		AssertEqual(t, value, 3)
		value, _ = myStackOInts.Pop()
		AssertEqual(t, value, 1)
		AssertTrue(t, myStackOInts.IsEmpty())
	})

	t.Run("stack of strings", func(t *testing.T) {
		myStackOStrings := new(Stack[string])

		//check stack empty
		AssertTrue(t, myStackOStrings.IsEmpty())

		// add thing, check not empty
		myStackOStrings.Push("1")
		AssertFalse(t, myStackOStrings.IsEmpty())

		// add thing pop back again
		myStackOStrings.Push("3")
		value, _ := myStackOStrings.Pop()
		AssertEqual(t, value, "3")
		value, _ = myStackOStrings.Pop()
		AssertEqual(t, value, "1")
		AssertTrue(t, myStackOStrings.IsEmpty())
	})
}

func AssertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("Got %+v but wanted %+v", got, want)
	}
}

func AssertNotEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("Got %+v but wanted %+v", got, want)
	}
}

func AssertTrue(t testing.TB, got bool) {
	t.Helper()
	if !got {
		t.Errorf("Got %+v wanted true", got)
	}
}

func AssertFalse(t testing.TB, got bool) {
	t.Helper()
	if got {
		t.Errorf("Got %+v want false", got)
	}
}
