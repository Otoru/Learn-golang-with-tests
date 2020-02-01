package main

import "math"

// Shape is implemented throughout the geometric shape
type Shape interface {
	Area() float64
}

// Rectangle is a geometric shape
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns the area of the geometric shape
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter returns the perimeter of the geometric shape
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Circle is a geometric shape
type Circle struct {
	Radius float64
}

// Area returns the area of the geometric shape
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle is a geometric shape
type Triangle struct {
	Base   float64
	Height float64
}

// Area returns the area of the geometric shape
func (c Triangle) Area() float64 {
	return (c.Base * c.Height) * 0.5
}

func main() {}
