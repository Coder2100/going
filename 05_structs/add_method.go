package main
import "fmt"

type Salary2 struct {
	Basic, HRA, TA float64
	breadth float64
	color string
}

type Employee3 struct {
	FirstName, LastName, Email string
	Age int
	MonthlySalary []Salary2

}

func (e Employee3) EmpInfo() string {
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)

	for _, info := range e.MonthlySalary{
		fmt.Println("=======================================================")
		fmt.Println(info.Basic)
		fmt.Println(info.HRA)
		fmt.Println(info.TA)
	}

	return "-------------------------------------------------------------"
}



func main(){
	e := Employee3{
		FirstName: "Zukile",
		LastName: "Mtotso",
		Email: "zmtotso.gmail.com",
		Age: 34,
		MonthlySalary: []Salary2{
			Salary2{
				Basic: 15000.00,
				HRA: 5000,
				TA: 1000,
			},
			Salary2{
				Basic: 15000.00,
				HRA: 5000.00,
				TA: 1000.00,
			},
			Salary2{
				Basic: 14000.00,
				HRA: 4500.00,
				TA: 2100,
			},
		},

	}
	fmt.Println(e.EmpInfo())
}