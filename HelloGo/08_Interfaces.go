package HelloGo

import (
	"math"
)

func Chapter_08_Interfaces() string {

	if ChapterIndex() &&
		Chapter_08_01_DeclaringAndImplementing() {
		return "O"
	}

	panic("OMG! Something is wrong in chapter 8!")
}

func Chapter_08_01_DeclaringAndImplementing() bool {
	// See code below the function for actual declaration and
	// implementation.

	// A generic function that, given a Geometry will return its area.
	calculateArea := func(geometry Geometry) float64 {
		return geometry.getArea()
	}

	// We can either a Circle or a Rectangle to calculateArea()
	// as they both implement the Geometry interface by defining
	// the function getArea()

	aCircle := Circle{1}

	if calculateArea(aCircle) != math.Pi {
		return false
	}

	aRectangle := Rectangle{10, 20}

	if calculateArea(aRectangle) != 200 {
		return false
	}

	return true
}

// Declare the interface.
type Geometry interface {
	getArea() float64
}

// In Go you don't explicitly implement an interface, a Circle
// or a Rectangle can be used as a Geometry because it implements getArea().
type Circle struct {
	radius float64
}

func (circle Circle) getArea() float64 {
	return circle.radius * circle.radius * math.Pi
}

type Rectangle struct {
	width  float64
	height float64
}

func (rectangle Rectangle) getArea() float64 {
	return rectangle.width * rectangle.height
}
