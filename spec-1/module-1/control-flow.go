package module1

import (
	"fmt"
)

func ControlFlow() {
	// if else
	ifElseDemo(5)
	ifElseDemo(6)
	ifElseDemo(7)
	conciseIf()

	// switch
	tagLessSwitch(1)
	tagLessSwitch(2)

	tagFullSwitch(1)
	tagFullSwitch(2)
	tagFullSwitch(3)

	// fmt.Print("hello, whats you name ? ")
	// var name string
	// _, err := fmt.Scan(&name)
	// if err == nil {
	// 	fmt.Println("Welcome " + name)

	// }

}

// if else
func ifElseDemo(x int) {
	if x == 5 {
		fmt.Println("x is 5")
	} else if x == 6 {
		fmt.Println("x is 6")
	} else {
		fmt.Println("x is neither 5 nor 6")
	}
}

func conciseIf() {
	name := "piyush"

	if n := len(name); n < 10 {
		fmt.Println("name length is less than 10")
	} else {
		fmt.Println("name length is not less than 10")
	}
}

/*
	In Go, the break statement is not required at the end of a case to stop it from falling through to the next case.
	The break statement is implicit in Go.

	If you do want a case to fall through to the next case, you can use the fallthrough keyword.
*/

// in tagLess switch, first condition matching the condition, will be executed!
func tagLessSwitch(x int) {
	switch {
	case x > 1:
		fmt.Println("Hello world from 1")
		fmt.Println("x is greater than 1")

	case x < 1:
		fmt.Println("hello from underground world")
		fmt.Println("x is less than 1")

	default:
		fmt.Println("i am one")
	}
}

func tagFullSwitch(x int) {
	switch x {
	case 1:
		fmt.Println("it is one")
	case 2:
		fmt.Println("it is 2")
	case 3:
		fmt.Println("it is other")
	}
}
