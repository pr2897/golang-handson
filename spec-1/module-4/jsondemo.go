package module4

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string
	Addr  string
	Phone string
}

func JsonDemo() {
	p1 := Student{Name: "Piyush", Addr: "mg road", Phone: "123"}
	baar, _ := json.Marshal(p1)

	fmt.Println(p1)
	fmt.Println(baar)

	var p2 Student
	err := json.Unmarshal(baar, &p2)
	fmt.Println(err)
	fmt.Println(p2)

}
