package main

import "fmt"

func main() {
	fmt.Println("Struc")
 
	userInfo := User{"golang", "golang@gmail.com", 23, false}

	fmt.Println( userInfo)
fmt.Println("user info %v\n", userInfo)
fmt.Printf("name is %v, email is %v and age %v.", userInfo.Name, userInfo.Email , userInfo.age)
}

type User struct{
	Name string
	Email string
	age int 
	status bool
}