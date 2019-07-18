package main

import "fmt"

type rectangle4 struct {
	length int
	breadth int
	color string
}

func main(){
	rect1 := new(rectangle4)//rect1 is a pointer to an instance of rectangle
	// new will enable us to use dot notation to assign values
	rect1.length = 10
	rect1.breadth = 20
	rect1.color = "Green"
	fmt.Println("Hi", rect1)

	var rect2 = new(rectangle4)// rect2 is an instance of rectangle4
	rect2.length= 10
	rect2.color = "Red"
	fmt.Println(rect2)
}