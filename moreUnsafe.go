//A Go program using map with sharding
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	array := [...]int{0, 1, -2, 3, 4}
	pointer := &array[0]
	fmt.Print(*pointer, " ")
	memoryAddress := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	//the pointer variable is converted to an unsafe.Pointer() and then to an uintptr
	for i := 0; i < len(array)-1; i++ {
		pointer = (*int)(unsafe.Pointer(memoryAddress))
		fmt.Print(*pointer, " ") //The *pointer notation dereferences the pointer
		//and returns the stored integer value.
		memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
		//The value of unsafe.Sizeof(array[0]) is what gets you to the next element of the array
		//because this is the memory occupied by each array element.
	}
	//we are trying to access an element of the array that does not exist using
	//pointers and memory addresses
	fmt.Println()
	pointer = (*int)(unsafe.Pointer(memoryAddress))
	fmt.Print("One more: ", *pointer, " ")
	memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	fmt.Println()
}
