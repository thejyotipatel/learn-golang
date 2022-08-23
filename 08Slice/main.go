package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slice")

	var sliceList = []string{"apple", "banana","mango"}

	fmt.Println("Slice :", sliceList)
	fmt.Printf("Slice type: %T\n", sliceList)
sliceList[0] = "guyuj"
	sliceList = append(sliceList, "orange")
	fmt.Println("Slice updated: ", sliceList)
	
	// slice the slice
	fmt.Println("Slice 1: - ", sliceList[1:])
	fmt.Println("Slice 1:3 - ", sliceList[1:3])
	fmt.Println("Slice :3 - ", sliceList[:3])
	fmt.Println("Slice :1 - ", sliceList[:1])

// use of make
	nums := make( []int, 3)


	nums[0] = 1
	nums[1] = 4
	nums[2] = 3
	fmt.Println("Nnmbers by using make :", nums)
	
	// nums[3] = 4 this give error

	// use of append 
		nums = append(nums, 6, 0) 
		fmt.Println("Numbers by using make :", nums)

		// sort method
		fmt.Println("Numbers IntsAreSorted method:", 	sort.IntsAreSorted(nums))

  sort.Ints(nums)
		fmt.Println("Numbers sort method:", nums)
	 
		fmt.Println("Numbers IntsAreSorted method:", 	sort.IntsAreSorted(nums))


	// how to remove a value from slice based on index

	var languagea = []string{"JS", "Swift", "CSS"}

	fmt.Println(languagea)

	var index = 1

	languagea = append(languagea[:index], languagea[index+1:]... )

	fmt.Println(languagea)
}