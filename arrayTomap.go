package main

import (
	"fmt"
)

func main() {
	anArray := [4]int{1, 2, 4, -4}
	arrayToMap := make(map[int]int)

	for _, s := range anArray {
		arrayToMap[s] = s

		fmt.Print(s, " ")

		fmt.Println()
	}
	fmt.Println("arrayToMap:", arrayToMap)

}
