package HelloGo

// This is the prototype of the function.
// See the .s file for the actual implementation.
func add(x, y int64) int64

func Chapter_99_Assembly() string {
	if add(22, 11) == 33 {
		return "D"
	}

	panic("OMG! Something is wrong in chapter 99!")
}
