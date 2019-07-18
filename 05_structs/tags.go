package main
import (
	"encoding/json"
	"fmt"
)

type Employee2 struct{
	FirstName string `json: "firstname"`
	LastName string  `json: "lastname"`
	City string `json: "city"`
}

func main(){
	json_string := `
{
"firstname": "String",
"lastname": "String",
"city": "London"
}`

emp1 := new(Employee2)
json.Unmarshal([]byte(json_string), emp1)
fmt.Println(emp1)
emp2 := new(Employee2)

emp2.FirstName = "Ramesh"
emp2.LastName = "Soni"
emp2.City = "Mumbai"
jsonStr, _ := json.Marshal(emp2)
fmt.Printf("%s\n", jsonStr)
}
