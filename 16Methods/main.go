package main

import "fmt"

func main() {
	fmt.Println("Meathod")
 
	userInfo := User{"golang", "golang@gmail.com", 23, false}

// 	fmt.Println( userInfo)
// fmt.Println("user info %v\n", userInfo)
fmt.Printf("name is %v, email is %v and age %v.", userInfo.Name, userInfo.Email , userInfo.age)

userInfo.GetName()

userInfo.newAge()
fmt.Printf("name is %v, email is %v and age %v.", userInfo.Name, userInfo.Email , userInfo.age)
}

type User struct{
	Name string
	Email string
	age int 
	status bool
}

func (u User) GetName() {
	fmt.Println("\nUser Name: ", u.Name)
}

func (u User) newAge()  {
	u.age = 22
	fmt.Println("New Age of this menthod: ", u.age)
}