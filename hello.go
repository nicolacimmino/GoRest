package main

import (
	"fmt"
	// This would be imported as HelloGo by default (name of the
	// last part of the import), we can give it another name.
	o "HelloGo/HelloGo"
)

func main() {
	fmt.Println(
		o.Index() +
			o.Chapter_01_Declarations() +
			o.Chapter_02_Functions() +
			o.Chapter_03_Types() +
			o.Chapter_04_ControlStructures() +
			o.Chapter_05_ArrayAndSlices() +
			o.Chapter_06_Maps() +
			o.Chapter_07_Structs() +
			o.Chapter_08_Interfaces() +
			o.Chapter_99_Assembly())
}
