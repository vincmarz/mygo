//Write a Go program that keeps reading integers until it gets the word END as
//input.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var f *os.File  //file pointer f
	f = os.Stdin    // file pointer declaration
	defer f.Close() // we defer the closing of that file with .Close()

	scanner := bufio.NewScanner(f) // return a new Scanner to read from f
	for scanner.Scan() {           // scanner is used with Scan() function for reading from os.Stdin line by line
		if scanner.Text() == "END" {
			fmt.Println("EXIT")
			os.Exit(1)
		} else {
			str, err := strconv.Atoi(scanner.Text())
			if err == nil {
				fmt.Println(str)
			} else {
				fmt.Println("Not integer number")
			}
		}

		//fmt.Println(">", scanner.Text())
	}
}
