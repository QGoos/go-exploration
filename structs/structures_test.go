package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := Perimeter(rect)
	want := 40.0
	if got != want {
		t.Errorf("expected %g but got %g", want, got)
	}
}

func TestArea(t *testing.T) {

	// sample of table driven tests
	// this does the same as the rectangles and circles combined
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{10.0}, 314.1592653589793},
	}

	for _, st := range areaTests {
		assertCorrectMessage(t, st.shape, st.want)
	}

	t.Run("rectangles", func(t *testing.T) {
		rect := Rectangle{12.0, 6.0}
		want := 72.0
		assertCorrectMessage(t, rect, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793
		assertCorrectMessage(t, circle, want)
	})

}

func assertCorrectMessage(t testing.TB, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()
	if got != want {
		t.Errorf("%#v expected %g but got %g", shape, want, got)
	}
}
