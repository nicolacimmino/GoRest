package main

import "fmt"
import (
	"HelloGo/HelloGo"
)

func main() {
	var helloWorld string
	helloWorld =
		HelloGo.GiveMeAnH() +
			HelloGo.GiveMeAnE() +
			HelloGo.GiveMeAnL() +
			HelloGo.GiveMeAnL() +
			"o world" +
			HelloGo.GiveMeAnExclamationMark()
	fmt.Println(helloWorld)
}
