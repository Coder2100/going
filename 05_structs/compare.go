package main

import "fmt"

type rectangle6 struct {
	length float64
	breadth float64
	color string
}

func main() {
	var rect1 = rectangle6{10, 20, "Green"}
	rect2 := rectangle6{length: 20, breadth: 10, color: "Red"}

	if rect1 == rect2 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	rect3 := new(rectangle6)
	var rect4 = &rectangle6{}

	if rect3 == rect4 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}