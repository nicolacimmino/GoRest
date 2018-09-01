package HelloGo

func Chapter_02_Functions() string {
	if true &&
		Chapter_02_01_Parameters() &&
		Chapter_02_02_VariadicFunctions() &&
		Chapter_02_03_MultipleReturnValues() &&
		Chapter_02_04_MultipleNamedReturnValues() &&
		Chapter_02_05_FunctionsInAVariable() &&
		Chapter_02_06_Closures() {
		return "E"
	}

	panic("OMG! Something is wrong in chapter 4!")
}

func Chapter_02_04_MultipleNamedReturnValues() bool {
	aFunction := func() (a rune, b int) {
		a = 'A'
		b = 4

		return
	}

	myA, myB := aFunction()

	if myA != 'A' || myB != 4 {
		return false
	}

	aAnoterFunction := func() (int, int) {
		return 12, 13
	}

	_, secondValue := aAnoterFunction()

	if secondValue != 13 {
		return false
	}

	return true
}

func Chapter_02_03_MultipleReturnValues() bool {
	aFunction := func() (int, int) {
		return 10, -10
	}

	x, y := aFunction()
	return x == 10 && y == -10
}

func Chapter_02_01_Parameters() bool {
	aFunction := func(a int) int {
		return a * 10
	}

	aValue := aFunction(20)

	return aValue == 200
}

func Chapter_02_05_FunctionsInAVariable() bool {
	sumFunction := func(a int, b int) int {
		return a + b
	}
	var sum = sumFunction(5, -5)

	return sum == 0
}

func Chapter_02_06_Closures() bool {
	localVariable := 10

	aFunctionThatExecutesAFunction := func(someFunction func() int) int {
		return someFunction()
	}

	if aFunctionThatExecutesAFunction(func() int {
		// localVariable is in scope also here
		return localVariable - 5
	}) != 5 {
		return false
	}

	aCounterBuilder := func() func() int {
		counter := 0
		return func() int {
			counter++
			return counter
		}
	}

	counterA := aCounterBuilder()
	counterB := aCounterBuilder()
	valA1 := counterA()
	valA2 := counterA()
	valB1 := counterB()

	if valA1 != 1 || valA2 != 2 || valB1 != 1 {
		return false
	}

	return true
}

func Chapter_02_02_VariadicFunctions() bool {
	aFunction := func(aValue int, otherValues ...int) int {
		grandTotal := aValue

		for _, value := range otherValues {
			grandTotal += value
		}

		return grandTotal
	}

	total := aFunction(10, 20, 30, 40, 50)

	return total == 150
}
