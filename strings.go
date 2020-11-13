package main

import (
	"fmt"
)

func main() {

	const sLiteral = "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x9c"
	fmt.Println(sLiteral)
	fmt.Printf("x: %x\n", sLiteral)

	fmt.Printf("sLiteral length: %d\n", len(sLiteral))

	for i := 0; i < len(sLiteral); i++ {
		fmt.Printf("%x ", sLiteral[i])
	}
	fmt.Println()

	fmt.Printf("q: %q\n", sLiteral)   // print a double-quoted string that is safely escaped with Go syntax
	fmt.Printf("+q: %+q\n", sLiteral) // print a string as ASCII-only output
	fmt.Printf(" x: % x\n", sLiteral) // put space between the printed bytes

	fmt.Printf("s: As a string: %s\n", sLiteral) //print a string literal as a string

	s2 := "€£³" // a string named s2 with three Unicode characters
	for x, y := range s2 {
		fmt.Printf("%#U starts at byte position %d\n", y, x) //print the characters in the U+0058 format
	}

	fmt.Printf("s2 length: %d\n", len(s2))

	const s3 = "ab12AB"
	fmt.Println("s3:", s3)
	fmt.Printf("x: % x\n", s3)

	fmt.Printf("s3 length: %d\n", len(s3))

	for i := 0; i < len(s3); i++ {
		fmt.Printf("%x ", s3[i])
	}
	fmt.Println()
}
