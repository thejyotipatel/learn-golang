package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time")

	sec := time.Now()
	// fmt.Printf("Time", sec)

	fmt.Println(sec.Format("01-02-2006 15:04:05 Monday"))

	createTime := time.Date(2020, time.August, 10, 23, 15, 0, 0, time.Local)
	fmt.Println(createTime)
}
