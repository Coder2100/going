package main

import (
	"fmt"
	"math"
)

//create const
// const has no type until its allocated
const name string = "Odwa"

func main() {
	fmt.Println(name)
	const n = 5000000
	const d = 3e20 / n
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
