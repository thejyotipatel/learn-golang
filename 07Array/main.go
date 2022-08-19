package main

import "fmt"

func main() {
	fmt.Println("Array")

	var language [4]string
	language[0] = "JS"
	language[1] = "HTML"
	language[3] = "CSS"

	fmt.Println("Language :", language)
	fmt.Println("Language len :", len(language))

	var num = [3]int {1,2,3}
	fmt.Println("Numbers :", num)
	fmt.Println("Numbers :", len(num)) 
}  