package module3

import "fmt"

type Student struct {
	Name string
	Age  int
}

func StructDemo() {

	piyush := Student{Name: "Piyush", Age: 23}
	fmt.Println(piyush)
}
