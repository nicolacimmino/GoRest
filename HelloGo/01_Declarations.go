package HelloGo

func Chapter_01_Declarations() string {
	// Declare and assign a value
	var asciiA int32
	asciiA = 65

	// Declare and initialise
	var aToHAsciiOffset float32 = 7

	// Declare and initialise, type is inferred, int in this case
	var zeroOffset int32 = 0

	// Declare and initialise multiple variables
	var uselessValueA, uselessValueB, uselessValueC int32 = 10, -3, -7

	// Declare and initialise shorthand, valid only inside functions, type always inferred
	anotherZero := int32(0)

	// Do the math!
	var asciiH = asciiA + int32(aToHAsciiOffset) + zeroOffset + uselessValueA + uselessValueB + uselessValueC + anotherZero

	if asciiH != rune('H') {
		panic("OMG! Something is wrong in chapter 1!")
	}

	return string(asciiH)
}

func Index() string {
	return ""
}
