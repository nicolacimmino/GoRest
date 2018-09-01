package HelloGo

func Chapter_04_ControlStructures() string {

	if Cahpter_04_01_If() &&
		Chapeter_04_02_Loops() &&
		Chapeter_04_03_Switch() {
		return "L"
	}

	panic("OMG! Something is wrong in chapter 4!")
}

func Cahpter_04_01_If() bool {
	anInt := 32

	if anInt != 32 {
		return false
	} else {
		return true
	}

	// You can have an expression before the if. If you define a variable
	// here it will be in scope only of the if clause and body.
	if aBiggerInt := anInt + 10; aBiggerInt != 42 {
		return false
	}
	// Cannot use aBiggerInt here.

	// This works also with multiple values, for instance we can type assert
	// and take action only if the value is of a certain type.
	var aValue interface{} = int(12)
	if _, isInt := aValue.(int); !isInt {
		return false
	}

	return true
}

func Chapeter_04_02_Loops() bool {

	// A classic for
	aValue := 0
	for index := 0; index < 10; index++ {
		aValue++
	}

	// The equivalent of a while in other languages
	for aValue < 11 {
		aValue++
	}

	// The equivalent a while(true) in other languages
	for {
		aValue++
		break
	}

	// We can break out of nested loops by defining a label
	// and break on that.
myLoop:
	for {
		for {
			aValue++
			break myLoop
		}
	}

	if aValue != 13 {
		return false
	}

	return true
}

func Chapeter_04_03_Switch() bool {
	aValue := 145

	// A standard switch
	switch aValue {
	case 145:
		aValue++
		// Cases don't fall through, we don't need a break here.
	case 200, 404, 500: // You can have multiple values for each case.
		return false
	default:
		return false
	}
	if aValue != 146 {
		return false
	}

	// A switch with conditions on the case.
	switch {
	case aValue > 100:
		aValue++
	case aValue == 101:
		return false
	default:
		return false
	}
	if aValue != 147 {
		return false
	}

	// You can declare a value in the switch which will be in scope
	// only of the switch condition and body.
	switch doubleAValue := aValue * 2; {
	case doubleAValue > 200:
		aValue++
	default:
		return false
	}
	// Cannot use doubleAValue here
	if aValue != 148 {
		return false
	}

	return true
}
