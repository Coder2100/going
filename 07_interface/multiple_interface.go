package main

import "fmt"

type Polygons2 interface {
	Perimeter2()
}

type object interface {
	NumberOfSide()
}

type Pentagon2 int

func (p Pentagon2) Perimeter2(){
	fmt.Println(5 *p)
}

func (p Pentagon2) NumberOfSide(){
	fmt.Println("Pentagon has 5 sides.")
}

func main(){
	obj := Pentagon2(50)
	obj.Perimeter2()
	obj.NumberOfSide()

}