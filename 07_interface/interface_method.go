package main
/* The key feature about interfaces is that interfaces are an abstract type,
they are described by their behavior (via their methods), and the types that -
satisfy the interface may implement the methods. The methods of an interface-
contain no code, they are abstract.*/

import "fmt"

type Vehicle interface {
	Structure()
	Speed()
}
type Human interface {
	Structure()
	Performance()
}

type Car string
func (c Car) Structure() []string{
	var parts = []string{"ECU", "Engine", "Air Filters", "Wipers", "Gas Task"}
	return parts
}

type Man string

func (m Man) Structure() []string {
	var parts = []string{"Brain", "Heart", "Nose","Bums", "Toes", "Teaths"}
	return parts
}

func main(){
	var c Car
	var m Man

	for i, j := range c.Structure(){
		fmt.Printf("%-15s <-----------------------------------> %15s\n", j, m.Structure()[i])
	}
}