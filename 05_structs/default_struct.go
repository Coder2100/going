package main


import "fmt"

type Employee4 struct{
	Name string
	Age int
}

func (obj *Employee4) Info(){
	if obj.Name == ""{
		obj.Name = "Zukile Mtotso"
	}
	if obj.Age == 0 {
		obj.Age = 25
	}
	}

	func main(){
		emp1 := Employee4{Name: "Zukile Mtotso"}
		emp1.Info()
		fmt.Println(emp1)

		emp2 := Employee4{Age: 26}
		emp2.Info()
		fmt.Println(emp2)

	}
