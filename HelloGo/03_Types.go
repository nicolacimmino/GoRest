package HelloGo

import (
	"math"
	"unicode/utf8"
	"unsafe"
)

/******************
Types
*/
func GiveMeAnL() string {

	if convertingIntToSmallerIntTruncates() &&
		convertingFloatToSmallerIntTruncates() &&
		isRuneAliasForInt32() &&
		isBoolWorking() &&
		isByteWorking() &&
		doIntegersWork() &&
		doUnsignedIntegersWork() &&
		doPointersWork() &&
		doStringsWork() &&
		doFloatsWork() &&
		doTypeAssertionsWork() {
		return "L"
	}

	return ""
}

/**
 * 03.1 bool
 *
 * Boolean can have only "true" or "false" value. There is no implicit conversion
 * from 0, 1, empty string etc as in some other languages.
 */
func isBoolWorking() bool {
	b := true

	// You can apply the negate operator to a bool
	b = !b

	if b {
		return false
	}

	return true
}

/**
 * 03.2 byte
 *
 * The byte type is an alias for uint8.
 */
func isByteWorking() bool {
	var b byte = 0

	// We can find the max value a byte can hold by wrapping it around below zero,
	// if it is really a uint8 this will give 255.
	maxByteValue := b - 1

	// Another proof this is a uint8, adding any other type to it would fail at compile time.
	b = b + uint8(2)

	return maxByteValue == 255 && b == 2
}

/**
 * 03.3 int int8 int16 int32 int64
 *
 * All the intX types are signed integers.
 */
func doIntegersWork() bool {
	// How big is an int? It depends on the architecture your program is being
	// compiled for. It will be 32-bit on a 32-bit processor and 64 on a 64-bit
	// processor, so you'll need to be careful if the values you are storing
	// could exceed 2^32.
	var anInt int = -1
	if unsafe.Sizeof(anInt) != 4 && unsafe.Sizeof(anInt) != 8 {
		return false
	}

	// Get byte size of an intX. This will be 1,2,4,8 respectively for int8,16,32,64.
	var anInt8 int8 = -1
	if unsafe.Sizeof(anInt8) != 1 {
		return false
	}

	return true
}

/**
 * 03.4 uint uint8 uint16 uint32 uint64
 *
 * All the uintX types are un-signed integers.
 */
func doUnsignedIntegersWork() bool {
	// How big is a uint? It depends on the architecture your program is being
	// compiled for. It will be 32-bit on a 32-bit processor and 64 on a 64-bit
	// processor, so you'll need to be careful if the values you are storing
	// could exceed 2^32.
	var aUInt uint = 1
	if unsafe.Sizeof(aUInt) != 4 && unsafe.Sizeof(aUInt) != 8 {
		return false
	}

	// Get byte size of an intX. This will be 1,2,4,8 respectively for uint8,16,32,64.
	var aUInt8 uint8 = 1
	if unsafe.Sizeof(aUInt8) != 1 {
		return false
	}

	// Overflowing an unsigned will wrap around the value
	aUInt8 = 0
	aUInt8 = aUInt8 - 1
	if aUInt8 != 255 {
		return false
	}

	return true
}

/**
 * 03.5 uintptr
 *
 * Note there is no pointer arithmetic in Go.
 */
func doPointersWork() bool {
	var anInt32 int32 = 0

	// int32Pointer will be of type uintptr, note this doesn't mean that it points
	// to a uint (it points to an int32) but that the pointer a uint.
	int32Pointer := &anInt32

	// You can dereference a point to get to the value with the * operator.
	// This works both with L and R values.
	*int32Pointer = *int32Pointer + 10

	// Above we added 10 to anInt32
	if anInt32 != 10 {
		return false
	}

	return true
}

func doFloatsWork() bool {
	var aFloat32 float32 = 1.1

	twoFloat := float32(2)
	if (aFloat32 * twoFloat) != 2.2 {
		return false
	}

	// You cannot multiply a float by an int, which conveniently,
	// removes the ambiguity on whether a float*int will return an
	// int or a float. You will need to convert the type of one or
	// the other. Here the multiplication is done between floats and
	// will result in 3.3000002, we accept anything with error
	// up to 0.1, just for this example this is not necessarily a
	// good way to compare floats depending on the application.
	// Note that math functions require float64s.
	threeInt := int(3)
	if math.Abs(float64(aFloat32*float32(threeInt))-3.3) > 0.1 {
		return false
	}

	// In this case we convert the float to int so we are really
	// computing 3*1 and the result will be int 3.
	if (int(aFloat32) * threeInt) != 3 {
		return false
	}

	// As above for a float64
	aFloat64 := float64(2.2)
	if math.Abs((aFloat64*float64(aFloat32))-2.42) > 0.1 {
		return false
	}

	return true
}

/**
 * The rune type is just an alias for int32 and can contain a single unicode character.
 */
func isRuneAliasForInt32() bool {
	// As a literal a rune is enclosed in single quotes.
	var aRune = 'A'

	// We can do math on a rune as if it was an int32 (because it really is)
	// For instance we convert the uppercase A to lowercase adding an offset of 32.
	aRune = aRune + 32

	if aRune == 'a' && aRune == 97 {
		return true
	}

	return false
}

/**
 * 03.6 string
 */
func doStringsWork() bool {
	// A string literal is enclosed in double quotes.
	var aString = "語Test"

	// Strings are concatenated using the + operator
	aString = aString + "!"

	// A string is just a sequence of bytes. The len() function will return
	// a bytes count and not the characters length. The string "語Test!" has
	// len() 8 because the first symbol takes 3 bytes in UTF-8.
	if len(aString) != 8 {
		return false
	}

	// To correctly get a count the characters in a string you can convert
	// it first to an array of runes. "語Test!" has 6 characters.
	if len([]rune(aString)) != 6 {
		return false
	}

	// Alternatively you can use the RuneCountInString in the utf8 package.
	if utf8.RuneCountInString(aString) != 6 {
		return false
	}

	// A string is just a sequence of bytes you can access as an array,
	// it is NOT a sequence of characters. So expecting aString[1] to contain
	// a T would happen to work only if the preceding character was not longer
	// than 1 byte in UTF-8.
	if aString[0] != '\xe8' || // UTF-8 for 語 is \xe8\xaa\x9e
		aString[1] != '\xaa' ||
		aString[2] != '\x9e' ||
		aString[3] != 'T' {
		return false
	}

	// To safely get the n-th character you will need to convert the string
	// to an array of runes and then take the n-th rune. "語Test!" second
	// character is a T.
	if string([]rune(aString)[1]) != "T" {
		return false
	}

	// In the same way you can safely get a substring by slicing the array
	// of runes. Get everything from "語Test!" except the first character.
	if string([]rune(aString)[1:]) != "Test!" {
		return false
	}

	// Range works correctly with UTF-8 strings, so the second character
	// of "語Test!" is actually a T.
	for index, runeValue := range aString {
		if index == 1 && runeValue != 'T' {
			return false
		}
	}

	// You can escape special chars in a string literal.
	aTwoLineString := "line1\nline2"
	if []rune(aTwoLineString)[5] != 10 {
		return false
	}

	// Or you can have a string literal that spans multiple lines
	// using back-ticks, unfortunately you cannot escape back-ticks
	// inside a back-tick string, so if you need to insert one you
	// will have to break out of the back-tick and concatenate.
	aMultipleLinesString :=
		`line1
		 line2 can contain a ` + "`" + `
		 line3
	`
	if []rune(aMultipleLinesString)[5] != 10 {
		return false
	}

	return true
}

func convertingIntToSmallerIntTruncates() bool {
	// Converting an int65 to an int8 truncates the higher bits
	var anInteger = int64(320)
	var an8BitInteger = int8(anInteger)
	return an8BitInteger == (320 & 255)
}

func convertingFloatToSmallerIntTruncates() bool {
	// Converting a float32 to a int8 converts to int and truncates the higher bits
	var aFloat32 = float32(320.32)
	var another8BitInteger = int8(aFloat32)
	return another8BitInteger == (320 & 255)
}

func doTypeAssertionsWork() bool {

	var aValue interface{} = "Test"

	// If aValue was not a string this would cause a panic.
	theUnderlyingString := aValue.(string)
	if theUnderlyingString != "Test" {
		return false
	}

	// We can check the underlying value type without causing
	// a panic by getting the second return value and checking for it.
	_, isInt := aValue.(string)
	if !isInt {
		return false
	}

	return true
}
