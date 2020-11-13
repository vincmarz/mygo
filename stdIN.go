package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var f *os.File //file pointer f
	f = os.Stdin  // file pointer declaration
	defer f.Close() // we defer the closing of that file with .Close()

	scanner := bufio.NewScanner(f) // return a new Scanner to read from f
	for scanner.Scan() {  // scanner is used with Scan() function for reading from os.Stdin line by line
		fmt.Println(">", scanner.Text())
	}
}