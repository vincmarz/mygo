package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type myElement struct {
	Name    string
	Surname string
	Id      string
}

var DATA = make(map[string]myElement) //DATA is a map defined as global variable,
//keys are strings and values are myElement

func ADD(k string, n myElement) bool {
	if k == "" {
		return false
	}

	if LOOKUP(k) == nil {
		DATA[k] = n
		return true
	}
	return false
}

func DELETE(k string) bool {
	if LOOKUP(k) != nil {
		delete(DATA, k) //delete function for a map
		return true
	}
	return false
}

func LOOKUP(k string) *myElement {
	_, ok := DATA[k] //allow to determine whether a given key is in the map or not
	if ok {          //if the key exists
		n := DATA[k]
		return &n //return the myElement address
	} else {
		return nil
	}
}

func CHANGE(k string, n myElement) bool {
	DATA[k] = n
	return true
}

func PRINT() {
	for k, d := range DATA {
		fmt.Printf("key: %s value: %v\n", k, d)
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin) //use os.Stdin to read from the standard input stream
	for scanner.Scan() {                  //Scan advances the Scanner to the next token
		text := scanner.Text()         //return the most recent token generated by a call to Scan
		text = strings.TrimSpace(text) //returns a slice of the string with all leading and trailing white space removed
		tokens := strings.Fields(text) //splits the string around each instance of one or more consecutive white space characters

		fmt.Printf("%v", tokens)
		fmt.Println("<-- tokens")
		fmt.Printf("%v", len(tokens))
		fmt.Println("<-- len(tokens)")
		switch len(tokens) {
		case 0:
			continue
		case 1:
			tokens = append(tokens, " ")
			tokens = append(tokens, " ")
			tokens = append(tokens, " ")
			tokens = append(tokens, " ")
		case 2:
			tokens = append(tokens, " ")
			tokens = append(tokens, " ")
			tokens = append(tokens, " ")
		case 3:
			tokens = append(tokens, " ")
			tokens = append(tokens, " ")
		case 4:
			tokens = append(tokens, " ")
		}
		fmt.Printf("%v", tokens)
		fmt.Println("<-- tokens")
		fmt.Printf("%v", len(tokens))
		fmt.Println("<-- len(tokens)")

		// switch tokens[0] {
		// case "PRINT":
		// 	PRINT()
		// case "STOP":
		// 	return
		// case "DELETE":
		// 	if !DELETE(tokens[1]) {
		// 		fmt.Println("Delete operation failed!")
		// 	}
		// case "ADD":
		// 	n := myElement{tokens[2], tokens[3], tokens[4]}
		// 	if !ADD(tokens[1], n) {
		// 		fmt.Println("Add operation failed!")
		// 	}
		// case "LOOKUP":
		// 	n := LOOKUP(tokens[1])
		// 	if n != nil {
		// 		fmt.Printf("%v\n", *n)
		// 	}
		// case "CHANGE":
		// 	n := myElement{tokens[2], tokens[3], tokens[4]}
		// 	if !CHANGE(tokens[1], n) {
		// 		fmt.Println("Update operation failed!")
		// 	}
		// default:
		// 	fmt.Println("Unknown command - please try again!")
		// }
	}
}
