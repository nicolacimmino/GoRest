package HelloGo

import (
	"HelloGo/HelloGo/Vehicles"
)

func Chapter_07_Structs() string {

	if ChapterIndex() &&
		Chapter_07_01_DeclaringAndCreating() &&
		Chapter_07_02_Methods() &&
		Chapter_07_03_AnonymousStructs() &&
		Chapter_07_04_ConstructorMethods() &&
		Chapter_07_05_Embedding() {
		return "W"
	}

	panic("OMG! Something is wrong in chapter 7!")
}

func Chapter_07_01_DeclaringAndCreating() bool {
	// A struct is a type that is a collection of named values.
	type point struct {
		x float64
		y float64
	}

	// Create an instance.
	aPoint := point{65.1, 25.2}

	// Access members.
	if aPoint.x < 65 || aPoint.y < 25 {
		return false
	}

	return true
}

func Chapter_07_02_Methods() bool {
	aGpsPoint := gpsPoint{65.011528, 25.467091}

	// See below how isEastOfGreenwich() is attached to the struct,
	// this cannot be done inside a function.
	return aGpsPoint.isEastOfGreenwich()
}

type gpsPoint struct {
	lat float64
	lng float64
}

func (p gpsPoint) isEastOfGreenwich() bool {
	return p.lng > 0
}

func Chapter_07_03_AnonymousStructs() bool {
	// You can create an instance of a struct without creating a named struct type.
	point := struct {
		x float64
		y float64
	}{12.33, 45.9}

	return point.x > 12 && point.y > 45
}

func Chapter_07_04_ConstructorMethods() bool {
	// See the Vehicle package for the actual declarations.

	// There are no automatic constructors in Go, it's convention
	// to have an exported method NewType() to act as the constructor,
	// this needs to be called explicitly.
	bicycle := Vehicles.NewBicycle()

	// It's also good practice not to export the type so your package
	// user cannot just create an instance skipping the constructor method.
	// The line below wouldn't compile.
	//bicycle := Vehicles.bicycle{}

	if bicycle == nil {
		return false
	}

	return true
}

func Chapter_07_05_Embedding() bool {
	// See the Vehicle package for the actual declarations.

	// There are no automatic constructors in Go, it's convention
	// to have an exported method NewType() to act as the constructor,
	// this needs to be called explicitly.
	bicycle := Vehicles.NewBicycle()

	// The GetStartMethod() method exists on both bicycle and car but
	// it's implemented only once in abstract vehicle and made available
	// to both structs through embedding. See the Vehicles package.
	if bicycle.GetStartMethod() != "Pedal" {
		return false
	}

	car := Vehicles.NewCar()

	if car.GetStartMethod() != "Ignition key" {
		return false
	}

	return true
}
