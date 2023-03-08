package main

import (
	"fmt"
)

func main() {
	var i = 21
	fmt.Printf("%v \n", i)
	fmt.Printf("%T \n", i)
	fmt.Println("%")

	var j = true
	fmt.Printf("%t \n", j)
	fmt.Printf("%v \n", j)

	var value1 = 21
	fmt.Printf("%b \n", value1)

	var a int = 15
	fmt.Printf(" %x\n", a)

	var b int = 15
	fmt.Printf(" %X\n", b)

	var c string = "Я"
	fmt.Printf(" %X\n", c)

	var d float64 = 123.456
	fmt.Printf("%f\n", d)
	fmt.Printf("%e\n", d)

}
