package main

import "fmt"

func printType(i interface{}){
	fmt.Println(i)
}

func main(){
	var manyType interface{}
	manyType = 100
	fmt.Println(manyType)

	manyType = 200.50
	fmt.Println(manyType)

	manyType = "Xhosa"
	fmt.Println(manyType)

	printType("Go programming language")
	var countries  = []string{"india", "japan", "canada"}
	printType(countries)

	var employee = map[string]int{"Mark": 10, "Sandy": 20}
	//var employee = map[string]int{"Mark": 10, "Sandy": 20}
	printType(employee)

	country := [3]string{"Zimbabwe", "Lesotho", "Botswane"}
	printType(country)

/*
   The manyType variable is declared to be of the type interface{} and
it is able to be assigned values of different types.
The printType() function takes a parameter of the type interface{},
hence this function can take the values of any valid type
*/
}