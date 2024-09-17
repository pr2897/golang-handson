package module1

import "fmt"

func Pointers() {
	var x int = 1 // variable to type int

	var ip *int // pointer of type int

	ip = &x // points to address of x

	var y **int // pointer of type pointer of type int
	y = &ip     // points to address of ip

	fmt.Println(y)   // address of ip
	fmt.Println(*y)  // value at y i.e., value of ip.
	fmt.Println(ip)  // value of ip
	fmt.Println(**y) // value of x because of two derefering | y -> ip -> x
}
