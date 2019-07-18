//type assertion is a way to retrieve dynamic value from interface type

package main

import (
	"fmt"
	"reflect"
)

func main(){
	var anyType interface{}
	anyType = "Canada"
	fmt.Println("Variable type: ", reflect.TypeOf(anyType))
	str, ok := anyType.(string)//a type assertion form
	if ok {
		fmt.Println("Variables type: ", reflect.TypeOf(str))
	}else{
		fmt.Println("Variable is not string.")
	}

	var intType = 100
	anyType = intType
	fmt.Println("Variable type: ", reflect.TypeOf(anyType))


	var anotherType interface{}
	anotherType = anyType
	fmt.Println("Variable type: ", reflect.TypeOf(anotherType))
}