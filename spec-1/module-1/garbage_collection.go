package module1

import "fmt"

func Garbage_collection() {
	resp := test()
	fmt.Println(*resp)
}

func test() *int {
	x := 1
	return &x
}
