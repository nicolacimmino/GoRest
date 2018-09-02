package HelloGo

func Chapter_06_Maps() string {

	if Chapter_06_01_Maps() {
		return " "
	}

	panic("OMG! Something is wrong in chapter 6!")
}

func Chapter_06_01_Maps() bool {
	// Declare a map
	aMap := make(map[string]int)

	// Assign values by keys
	aMap["one"] = 1
	aMap["two"] = 2

	// Get a value by key
	value := aMap["two"]
	if value != 2 {
		return false
	}

	// If a key might not exist you can attempt to retrieve the value without
	// causing a panic. Value will default to zero in this case.
	value, indexPresent := aMap["three"]
	if indexPresent || value != 0 {
		return false
	}

	// Or you can just check if the index is present.
	if _, indexPresent := aMap["three"]; indexPresent {
		return false
	}

	// You can delete an element by key.
	delete(aMap, "two")
	if _, indexPresent := aMap["two"]; indexPresent {
		return false
	}

	// Declare a map as a literal
	anotherMap := map[string]int{
		"one": 1,
		"two": 2,
	}
	if len(anotherMap) != 2 {
		return false
	}

	// You can loop over a map
	for key, value := range anotherMap {
		if key == "two" && value != 2 {
			return false
		}
	}

	// You can use just the key or the value
	for _, value := range anotherMap {
		if value > 2 {
			return false
		}
	}

	return true
}
