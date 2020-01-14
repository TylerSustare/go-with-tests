package shapes

import "math"

// Shape is a shape interface
// In Go interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile.
type Shape interface {
	Area() float64
}

// Rectangle represents a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates area of a circle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle is a circle
type Circle struct {
	Radius float64
}

// Area calculates the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle is a triangle
type Triangle struct {
	Base   float64
	Height float64
}

// Area calculates the area of a triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// Perimeter calculates the perimeter
func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

// Area calculates area
func Area(r Rectangle) float64 {
	return r.Width * r.Height
}
