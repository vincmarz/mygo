package main

import (
	"fmt"
)

type Day int

func main() {

	const (
		Sunday Day = iota
		Monday
		Tuesday
		Wedneday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday)
	fmt.Println(Saturday)
}
