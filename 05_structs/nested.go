package main

import "fmt"

type Salary struct{
	Basic, HRA, TA float64
}

type Employee struct {
	FirstName, LastName, Email string
	Age int
	MonthlySalary []Salary
}

func main(){

	e := Employee{
		FirstName: "Jon",
		LastName: "Doe",
		Email: "js@foe.com",
		Age: 25,
		MonthlySalary:: []Salary{
			Basic: 15000.00,
			HRA: 5000.00,
			TA: 2000,
	},
	Salary{
		Basic: 16000,
		HRA: 5000,
		TA: 2100,
	},
	Salary{
		Basic: 17000,
		HRA: 5000,
		TA: 1000,
	},
	},

	//fmt.Printlne
}