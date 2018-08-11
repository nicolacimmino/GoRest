package HelloGo

/******************
Functions
*/
func GiveMeAnE() string {
	var baseChar, offset = giveMeECoordinates()
	var uselessOffsetA, uselessOffsetB = giveMeCancellingOffsets()
	var zero = multiplyByZero(123)

	// Function as variable value
	sumFunction := func(a int, b int) int {
		return a + b
	}
	var calculatedZero = sumFunction(5, -5)

	functionToCalculateAFive := giveMeAFunctionThatCalculatesAFive()
	five := functionToCalculateAFive()

	// Closures
	addOneFunction := func(a int) int {
		return a + 1
	}
	ten := runMyFunctionOnAValue(addOneFunction, 9)

	one := addALotOfNumbers(1, 10, 15, -15, -10)

	return string(int(baseChar) +
		offset +
		uselessOffsetA +
		uselessOffsetB +
		zero +
		calculatedZero +
		(five - 5) +
		(ten - 10) +
		(one - 1))
}

/**
Return multiple named parameters
*/
func giveMeECoordinates() (baseChar rune, offset int) {
	baseChar = 'A'
	offset = 4

	return
}

/**
Return multiple parameters
*/
func giveMeCancellingOffsets() (int, int) {
	return 10, -10
}

/**
Function parameters
*/
func multiplyByZero(value int) int {
	return value * 0
}

/**
Return a function.
Note that functionScopedVariable is in scope in the closure and will
be even when the result of giveMeAFunctionThatCalculatesAFive will be called
from our caller.
*/
func giveMeAFunctionThatCalculatesAFive() func() int {
	functionScopedVariable := 10

	// Function as variable value
	aClosure := func() int {
		// functionScopedVariable is in scope also here
		return functionScopedVariable - 5
	}

	return aClosure
}

/**
Closures
*/
func runMyFunctionOnAValue(f func(int) int, a int) int {
	return f(a)
}

/**
Variadic function
*/
func addALotOfNumbers(arguments ...int) int {
	sum := 0

	// rage iterates over an array and returns both the index and the value
	// we can use _ to ignore the index in this case.
	for _, value := range arguments {
		sum += value
	}

	return sum
}
