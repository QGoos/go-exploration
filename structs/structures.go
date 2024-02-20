package structs

import (
	"math"
)

// Shape Interface
type Shape interface {
	Area() float64
}

// Shape: Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Returns
// Area of Shape Rectangle
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Shape: Circle
type Circle struct {
	Radius float64
}

// Returns
// Area of Shape Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Accepts: rect Rectangle
// Returns: float64 perimeter of Rectangle
func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Height + rect.Width)
}

// Accepts: rect Rectangle
// Returns: float64 area of Rectangle
func Area(rect Rectangle) float64 {
	return rect.Height * rect.Width
}
