//Write a Go program that finds the sum of all command-line arguments that are
//valid numbers.
package main 

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more numbers")
		os.Exit(1)
	}	

	arguments := os.Args
    
    var sum float64 

    for i :=1; i < len(arguments); i++ {
    	char, err := strconv.ParseFloat(arguments[i], 64)
    	if err !=nil {
    		//fmt.Println(arguments[i], " is not a numbers")
    	} else {
            sum += char 
    		
    	}
    }

    fmt.Println("Sum: ", sum)
}