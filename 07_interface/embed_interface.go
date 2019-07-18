package main
/*Interfaces may embed other interfaces, this behavior is an aspect
of interface polymorphism which is known as ad hoc polymorphism.*/

/*When one type is embedded into another type, the methods of the embedded
type are available to the embedding type. The method or methods of the embedded
interface are accessible to the embedding interface.*/

import "fmt"

type Geometry4 interface {
	Edges() int
}

type Polygons4 interface {
	Geometry4//interface embedding another interface
}

type Pentagon4 struct{}
type Hexagon4 struct{}
type Octagon4 struct{}
type Decagon4 struct{}

func (p4 Pentagon4) Edges() int { return 5}
func (h4 Hexagon4) Edges() int { return 6}
func (o4 Octagon4) Edges() int { return 8}
func (d4 Decagon4) Edges() int { return 10}


func main(){
	p4 := new(Pentagon4)
	h4 := new(Hexagon4)
	o4 := new(Octagon4)
	d4 := new(Decagon4)

	polygons4 := [...]Polygons4{p4, h4, o4 ,d4}

	for i := range polygons4 {
		fmt.Println(polygons4[i].Edges())
	}
}
