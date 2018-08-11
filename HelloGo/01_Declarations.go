package HelloGo

/**
Declarations
*/
func GiveMeAnH() string {
	// Declare and assign a value
	var asciiA int
	asciiA = 65

	// Declare and initialise
	var aToHAsciiOffset float32 = 7

	// Declare and initialise, type is inferred, int in this case
	var zeroOffset = 0

	// Declare and initialise multiple variables
	var uselessValueA, uselessValueB, uselessValueC int = 10, -3, -7

	// Declare and initialise shorthand, valid only inside functions, type always inferred
	anotherZero := 0

	// Do the math!
	var asciiH = asciiA + int(aToHAsciiOffset) + zeroOffset + uselessValueA + uselessValueB + uselessValueC + anotherZero

	return string(asciiH)
}
