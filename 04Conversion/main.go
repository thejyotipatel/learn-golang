package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"   
	"strings"
)

func main() {
	fmt.Println("Enter values")

	reader := bufio.NewReader(os.Stdin)

	input , _ := reader.ReadString('\n')
	fmt.Println("thanks for values, ", input)

	values, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println("added 1 to value: ",values+1)
	}

}