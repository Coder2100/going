package main
import "fmt"

//Interfaces are types that declare sets of methods.

type MyInterface interface {
	methodone()// without parameter
	methodtwo(string)// with parameter
	methodthree()int // with return type

}

func main(){
	fmt.Print("Fuck the world")
}