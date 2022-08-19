package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("if else statement ")

	fmt.Println("Enter number to check is Even or Odd")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	num, err := strconv.Atoi(strings.TrimSpace(input))

	// err handler if else statement
	if err != nil{
		fmt.Println("It should be Number - ",err)
	}else{
		fmt.Println("number : ", num)
	}

	//  conditon in if else
	if num%2 == 0{
		fmt.Println("Even number")
	}else{
		fmt.Println("Odd number")
	}
}