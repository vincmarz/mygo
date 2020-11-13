//A Go program using C code
package main

//#include <stdio.h>
//void callC() {
//   printf("Calling C code!\n");
//}
import "C"

import "fmt"

func main() {
	fmt.Println("A Go statement!")
	C.callC() // in order to execute the callC() C function you will need to call it as C.callC()
	fmt.Println("Another Go statement!")

}
