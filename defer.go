package main

import (
	"fmt"
)

func p() {
	for i := 3; i > 0; i-- {
		fmt.Println(i)
	}
}

func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
}

func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}() //anonymous function takes no parameters
	}
	fmt.Println()
}

func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i) //anonymous function takes with one parameter
	}
}

func foo() (result string) {
	defer func() {
		//result = "Change World" // change value at the very last moment
		fmt.Println("Change World")
	}()
	return "Hello World"
}

func main() {
	p()
	d1()
	d2()
	fmt.Println()
	d3()
	fmt.Println()
	foo()
}
