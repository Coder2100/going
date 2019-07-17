package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")
	fmt.Println("1 +1 =", 1+1)

	fmt.Println("7.0/3.0 = ", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	//go variables are explicitly declared
	var a = "myfirst"
	var b, c int = 1, 3
	var d int
	f := "Zukile Mtotso"
	fmt.Println(a, b, c, d, f)
}
