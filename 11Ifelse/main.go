package main

import (
	"fmt"
)

// check even number
	func evnOdd(num int) string  {
	//  conditon in if else
	if num%2 == 0{
		return "Even number"
	} 
		return "Odd number" 

}

// find fib num using recurtion
func findFib(num int) (int) {
	 
if num  == 1 {
	return  1
}else {
	return findFib(num - 1) * num
}

	
}


// add all value using recursion
func addAll(num []int) int {
	 
if len(num)  == 0 {
	return 0
} 
rest  :=  num[1:] 
return num[0] + addAll(rest )  
}

func main() {
	fmt.Println("if else statement ")

	// reader := bufio.NewReader(os.Stdin)

	// input, _ := reader.ReadString('\n')

	// num, err := strconv.Atoi(strings.TrimSpace(input))

	// err handler if else statement
	// if err != nil{
	// 	fmt.Println("It should be Number - ",err)
	// }else{
	// 	fmt.Println("number : ", num)
	// }
	
	// result := evnOdd(num)
	// fmt.Println(result)

	// result := findFib(num)
	// fmt.Println(result)

 num := []int {1,2,-3,8,-5}
	result := addAll( num)
	fmt.Println(result)

}
 