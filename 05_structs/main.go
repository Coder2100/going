package main

import "fmt"

type user struct {
	username string
	email string
	password string
	date_of_birth string

}

//declration of struct typ
// the struct will not be imported to other
// package since the identifier starts with the lowercases

type rectangle struct{
	length float64
	breadth float64
	color string
}
func main(){
	fmt.Println("Hello world")
	fmt.Println(rectangle{10.5, 25.10, "yellow"})
	fmt.Println(user{"Zukile", "zmtotso@ymail.om", "password", "02April1991"})
}
