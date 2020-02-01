package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestPerimeter(t *testing.T) {

	assert := func(t *testing.T, result, expected float64) {
		t.Helper()
		if result != expected {
			t.Errorf("got '%f' and want '%f'.", result, expected)
		}
	}

	t.Run("Calculate perimeter", func(t *testing.T) {
		rectangle := Rectangle{5.0, 5.0}
		result := Perimeter(rectangle)
		expected := 20.0
		assert(t, result, expected)
	})

	t.Run("Calculate another perimeter", func(t *testing.T) {
		rectangle := Rectangle{7.5, 7.5}
		result := Perimeter(rectangle)
		expected := 30.0
		assert(t, result, expected)
	})
}

func ExamplePerimeter() {
	rectangle := Rectangle{5, 5}
	result := Perimeter(rectangle)
	fmt.Println(result)
	// Output: 20
}

func BenchmarkPerimeter(b *testing.B) {
	firstChoice := rand.Float64()
	secondChoice := rand.Float64()
	rectangle := Rectangle{firstChoice, secondChoice}
	for i := 0; i < b.N; i++ {
		Perimeter(rectangle)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		area  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 2, Height: 2}, area: 4},
		{name: "Circle", shape: Circle{Radius: 3}, area: 28.274333882308138},
		{name: "Triangle", shape: Triangle{Base: 2, Height: 2}, area: 2},
	}

	for _, item := range areaTests {
		t.Run(item.name, func(t *testing.T) {
			result := item.shape.Area()
			if result != item.area {
				t.Errorf("%#v got %g want %g", item.shape, result, item.area)
			}
		})

	}

}
