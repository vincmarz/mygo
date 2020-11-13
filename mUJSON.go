package main

import (
	"encoding/json"
	"fmt"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func main() {
	myRecord := Record{
		Name:    "Mihalis",
		Surname: "Tsoulkalos",
		Tel: []Telephone{{Mobile: true, Number: "1234-567"},
			{Mobile: true, Number: "1234-abcd"},
			{Mobile: true, Number: "abcc-567"},
		}}

	rec, err := json.Marshal(&myRecord)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(rec))
	//fmt.Println("Marshal")

	var unRec Record
	err1 := json.Unmarshal(rec, &unRec)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(unRec)
	//fmt.Println("Unmarshal")
}
