package module1

import (
	"fmt"
	"strconv"
	"strings"
)

func Datatypes() {
	// type aliasing
	type Interger int
	type Float float64

	// variable initialisation
	var age Interger
	var marks Float = 34.33

	age = 33

	// short hand declaration
	name := "piyush raj"

	fmt.Printf("%s %d %0.2f\n", name, age, marks)

	fmt.Println(strconv.ParseInt("11", 10, 0))
	fmt.Println(strings.ToUpper("hello world"))

	/*
		// using new keyword | alternate way to create a variable
		variable is initialised to zero
		new() creates a variable and returns the pointer to the variable.
	*/

	ptr := new(int)
	*ptr = 3
	fmt.Println(ptr, *ptr)

	fmt.Println("------------- variables end here ----------")
}
