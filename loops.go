package main

import (
	"fmt"
)

func main() {

	//for loop
	for i := 0; i < 100; i++ {
		if i%20 == 0 {
			continue
		}

		if i == 95 {
			break
		}

		fmt.Print(i, " ")
	}

	fmt.Println()
	i := 10
	//while loop
	for {
		if i < 0 {
			break
		}
		fmt.Print(i, " ")
		i--
	}

	fmt.Println()
	i = 0
	//do...while loop
	anExpression := true
	for ok := true; ok; ok = anExpression {
		if i > 10 {
			anExpression = false
		}

		fmt.Print(i, " ")
		i++
	}

	fmt.Println()
	anArray := [5]int{0, 1, -1, 2, -2}
	for i, value := range anArray {
		fmt.Println("index:", i, "value: ", value)
	}
	//Applying the range keyword to an array variable returns two values: an array index and
	//the value of the element at that index

}
