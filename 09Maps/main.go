package main

import "fmt"
 

func main() {
	fmt.Println("Maps")

	num := make(map[string]string)
	num["1"] = "One"
	num["2"] = "Two"
	num["4"] = "Four"

	fmt.Println("Numbers :" , num)
	fmt.Println("2 :" , num["2"])

	delete(num,"4")
	fmt.Println("Numbers :" , num)

	// loops in golang
	for key,value := range num{
		fmt.Printf("key %v, value %v\n", key,value)
	}
	
	for _,value := range num{

		fmt.Printf("value %v\n", value)
	}
}