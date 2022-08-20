package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch case")

	// get random number
	rand.Seed(time.Now().UnixNano())
	diceNum := rand.Intn(6) + 1

	fmt.Println("value :", diceNum)

	switch diceNum{
	case 1:
		fmt.Println("value is 1 and you can open")
	case 2:
		fmt.Println("you can move 2 spot")
	case 3:
		fmt.Println("you can move 3 spot")
		fallthrough
	case 4:
		fmt.Println("you can move 4 spot")
	case 5:
		fmt.Println("you can move 5 spot")
	case 6:
		fmt.Println("you can move 5 spot and roll agen")
	default:
		fmt.Println("try agen!")
	}
}