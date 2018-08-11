package main

import "fmt"
import "HelloGo/HelloGo"

func main() {
	var helloWorld string
	helloWorld = HelloGo.GiveMeAnH() +
		HelloGo.GiveMeAnE() +
		"llo world"
	fmt.Println(helloWorld)
}
