package  main

import "fmt"


type Salary struct {
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
		Email: "zmtotso@lat.com",
		Age: 15,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA: 4000,
				TA: 1000,
	      },
	     Salary{
	     	Basic: 16000,
	     	HRA: 4000,
	     	TA:1200,
	},
			Salary{
				Basic: 17000,
				HRA: 3000,
				TA: 1200,
			},
	},
	}
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	fmt.Println(e.MonthlySalary[0])
	fmt.Println(e.MonthlySalary[1])
	fmt.Println(e.MonthlySalary[2])
}