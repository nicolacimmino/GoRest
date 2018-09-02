package HelloGo

func Chapter_05_ArrayAndSlices() string {

	if ChapterIndex() &&
		Chapter_05_01_Arrays() &&
		Chapter_05_02_Slices() &&
		Chapter_05_03_Range() {
		return "O"
	}

	panic("OMG! Something is wrong in chapter 5!")
}

func Chapter_05_01_Arrays() bool {
	// An array of int with size 5
	var anArray [5]int
	anArray[0] = 10

	if anArray[0] != 10 {
		return false
	}

	// An array of size 3, only the first two elements initialised.
	// the 3rd element defaults to 0.
	var anInitialisedArray = [3]int{10, 20}

	if anInitialisedArray[1] != 20 {
		return false
	}

	// Compiler will figure out the size.
	var anotherInitialisedArray = [...]int{10, 20}
	if len(anotherInitialisedArray) != 2 {
		return false
	}

	return true
}

func Chapter_05_02_Slices() bool {
	someNumbers := []int{10, 20, 30, 40, 50, 60}

	// This is a slice of the array someNumbers. Note that aSlice
	// contains two elements (index 2 and 3 of someNumbers) as
	// the notation is low:high -> low inclusive, hi excluded.
	aSlice := someNumbers[2:4]

	// someNumbers[2] and someNumbers[3] are in aSlice[0] and aSlice[1]
	// respectively.
	if len(aSlice) != 2 || aSlice[0] != 30 || aSlice[1] != 40 {
		return false
	}

	// Or you can take a slice from index 2 to the end of the array
	anoterSlice := someNumbers[2:]

	if len(anoterSlice) != 4 {
		return false
	}

	// A slice is not a new array, it's not a copy of the values, it is
	// a reference to the original values.
	aSlice[0] += 1

	// So acting on aSlice[0] affected someNumbers[2]
	if someNumbers[2] != 31 {
		return false
	}

	// A slice of struct containing numbers and their english names.
	aSliceOfStructs := []struct {
		value       int
		englishName string
	}{
		{1, "one"},
		{2, "two"},
		{3, "three"},
		{4, "flour"},
	}
	// Note that 3 is the index in the slice, it has nothing to do with .value
	// This will replace the misspelled "flour" to "four" not "three"
	aSliceOfStructs[3].englishName = "four"

	if aSliceOfStructs[3].englishName != "four" {
		return false
	}

	return true
}

func Chapter_05_03_Range() bool {
	anArray := [10]int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}

	// You can loop over an array using range
	for index, value := range anArray {
		if index*10 != value {
			return false
		}
	}

	// You can do the same with just values if you
	// don't need the index.
	for _, value := range anArray {
		if value > 90 {
			return false
		}
	}

	// You can do same on a slice of the array. Here the
	// last 3 elements of anArray will be looped with the
	// slice index 0 to 2.
	for index, value := range anArray[7:] {
		if (index+7)*10 != value {
			return false
		}
	}

	// You can even discard both index and value, this will loop
	// 3 times as the slice size is 3.
	count := 0
	for range anArray[7:] {
		count++
	}

	if count != 3 {
		return false
	}

	return true
}
