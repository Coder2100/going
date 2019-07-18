//polymorphism is the one key functionality provided by the ineterfaces.
//using polymorpgic approach the function created, here PARAMETER is used by each Geometry types.

package main

import "fmt"

type Geometry interface {
	Edges() int
}

type Pentagon3 struct{}
type Hexagon3 struct{}
type Octagon struct{}
type Decagon struct{}

func (p3 Pentagon3) Edges() int { return 5 }
func (h3 Hexagon3) Edges() int  { return 6 }
func (o Octagon) Edges() int  { return 8 }
func (d Decagon) Edges() int  { return 10 }

func parameter(geo Geometry, value int) int {
	num := geo.Edges()
	calculation := num * value
	return calculation
}

func main() {
	p3 := new(Pentagon3)
	h3 := new(Hexagon3)
	o := new(Octagon)
	d := new(Decagon)

	g := [...]Geometry{p3, h3, o, d}

	for _, i := range g {
		fmt.Println(parameter(i, 5))
	}
}