//Write a Go program that finds the average value of floating-point numbers that
//are given as command-line arguments.
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
    var j float64 = 0 

    for i :=1; i < len(arguments); i++ {
    	char, err := strconv.ParseFloat(arguments[i], 64)
    	if err !=nil {
    		//fmt.Println(arguments[i], " is not a numbers")
    	} else {
    		j++
            sum += char 
    		
    	}
    }
    fmt.Println("average: ", sum / j)
}