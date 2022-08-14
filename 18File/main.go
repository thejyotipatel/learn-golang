package main

import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
)

func main() {
	fmt.Println("Files")
	contant := "Learning GOLang. -by golang"

	file, err := os.Create("./myfile.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, contant)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)


	fmt.Println("length is: ", length)
	defer file.Close()

	readFile("./myfile.txt")

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