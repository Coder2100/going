package main

import "fmt"

type rectange5 struct{
	length int
	breadth int
	color string
}


func main(){
	var rect1 = &rectange5{10, 20, "Green"}// cnnot skip any field
	fmt.Println(rect1)

	var rect2 = &rectange5{}
	// using dot notation to assign values
	rect2.length = 10
	rect2.color = "Orange"
	fmt.Println(rect2)

	var rect3 = &rectange5{}
	(*rect3).breadth = 10
	(*rect3).color = "Blue"

	fmt.Println(rect3)
}