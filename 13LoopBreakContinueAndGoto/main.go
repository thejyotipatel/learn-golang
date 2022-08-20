package main

import "fmt"

func main() {
	fmt.Println("Loops")

	days := []string{"sun", "mon","tue","wed","thu","fri","sat"}

	fmt.Println("Days: ",days)

	// type 1
	/* for i := 0; i<len(days); i++{
				fmt.Printf(days[i] + ", ")
	} */

	// type 2
	/* for i := range days{
		fmt.Printf(days[i] + ", ")
	} */
	
	// type 3
	/* for i, day := range days{
		fmt.Printf("index -%v day - %v\n",i,day)
	} */
 
	// type 4
	/* for _, day := range days{
		fmt.Printf("day- %v\n",day)
	} */

	
	// type 5
	day := 1

	 for day < 7{
			
			// use of break
			// if day == 4{
			// 	day++
			// 	break
			// }

			// use of continue
				if day == 4{
				day++
				fmt.Println("skip 4 and continue")
				continue
			}

			// use of goto statement
			if day == 2{
				goto gotoStatement
			}
		fmt.Printf("day- %v\n",day)
		day++
	} 

	// goto
	gotoStatement:
	fmt.Println("go to statement")
}