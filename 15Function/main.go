package main

import "fmt"

func main() {
	fmt.Println("Function")
 
 result := add(2 , 3)
	fmt.Println("Result: " , result)
 
 ProResult, msg := proAdd(2,3,4,6,3)
	fmt.Println("Pro Result: " , ProResult)
	fmt.Println("Pro Result message: ", msg)


}

// add two number
func add(num1 int, num2 int) int {
	return num1 + num2
}

//  add as many as you can number
func proAdd(num ...int)  (int,string) {
	total := 0

	for _, v := range num {
		total += v
	}

	return total, "pro added"
}