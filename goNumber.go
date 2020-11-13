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

	//if len(os.Args) == 2 {
	//	fmt.Println("One character")
	//}	

	arguments := os.Args
	//n, err := strconv.ParseFloat(arguments[1], 64)

	//fmt.Println(n)
    
    var sum float64 

    for i :=1; i < len(arguments); i++ {
    	char, err := strconv.ParseFloat(arguments[i], 64)
    	if err !=nil {
    		fmt.Println("Arguments not a numbers")
    		fmt.Println(char)
    	} else {
    		//sum, _ := strconv.ParseFloat(arguments[i], 64)
    		fmt.Println("sum:", sum)
            sum += char 
    		fmt.Println("sum + char: ", sum)
    	}
    }
    //sum := 0
     
    //fmt.Println("Sum: ", sum) 
    
	// if err !=nil {
	// 	fmt.Println("Arguments not a numbers")
	// 	os.Exit(2)	
	// } else {
	// 	for i :=1; i < len(arguments); i++ {
	// 		//sum := arguments[1] + arguments[i]
	// 		fmt.Println(i)
	// 	}
	// }

	//fmt.Println(sum)

}