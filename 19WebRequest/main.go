package main

import (
	"fmt"
	"io/ioutil"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web request")
	 
}

func readFile(fileName string)  {
	databyte, err := ioutil.ReadFile(fileName)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	fmt.Printf("file in bytes: ", string(databyte))
}

func checkNilError(err error)  {
	if err != nil {
		panic(err)
	}
}