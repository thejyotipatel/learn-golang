package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	user := "hello"
	fmt.Println(user)	

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the values:")

	// comma ok || error err
	input , _:= reader.ReadString('\n')
	fmt.Println("thanks for values, ", input)
fmt.Printf("Type of value %T",input)

}