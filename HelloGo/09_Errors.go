package HelloGo

import (
	"io/ioutil"
	"reflect"
)

func Chapter_09_Errors() string {

	if ChapterIndex() &&
		Chapter_09_01_ErrorHandling() {
		return "R"
	}

	panic("OMG! Something is wrong in chapter 9!")
}

func Chapter_09_01_ErrorHandling() bool {
	// There is no exception throw/catching, instead functions that can
	// fail will return an extra parameter that implements the error interface.
	_, err := ioutil.ReadFile("/tmp/afilethatreallyshouldntexist")

	if err == nil {
		return false
	}

	if _, isError := err.(error); !isError {
		return false
	}

	// You can get a string with the error message using the Error() method.
	message := err.Error()

	if reflect.TypeOf(message).Kind() != reflect.String {
		return false
	}

	return true
}
