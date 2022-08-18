package main

import "fmt"

func main() {
	fmt.Println("Pointer")

	// var ptr *int 
	// fmt.Println("value: ", ptr)

	num := 8

	var ptr  = &num
	fmt.Println("value ",ptr)
	fmt.Println("value ",*ptr)
	
	*ptr = *ptr + 2
	fmt.Println("value ",num)

}