package main
import "fmt"

type rectangle3 struct{
	length int
	breadth int
	color string

}



func main(){
	var rect1 = rectangle3{10, 20, "green"}
	fmt.Println(rect1)

	var rect2 = rectangle3{length:10, color:"Green"}
	fmt.Println(rect2)

	rect3 := rectangle3{10,20, "Green"}

	fmt.Println(rect3)

	rect4 := rectangle3{length:10, breadth: 20, color:"Green"}

	fmt.Println(rect4)
	rect5 := rectangle3{breadth: 20, color:"Green"}
	fmt.Println(rect5)
}