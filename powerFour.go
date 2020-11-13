package main

import (
	"fmt"
)

type Power4 int

const PI = 3.1415926

const (
	C1 = "C1C1C1"
	C2 = "C2C2C2"
	C3 = "C3C3C3"
)

func main() {

	const (
		p4_0 Power4 = 1 << 0 << iota
		p4_1
		p4_2
		p4_3
	)

	fmt.Println("4^0:", p4_0)
	fmt.Println("4^1:", p4_1)
	fmt.Println("4^2:", p4_2)
	fmt.Println("4^4:", p4_3)
}
