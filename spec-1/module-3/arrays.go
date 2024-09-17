package module3

import "fmt"

// fixed length

func ArraysDemo() {
	var x [5]int

	x[0] = 1
	fmt.Println(x)

	predefined := [5]int{1, 2, 3, 4, 5} // we defined it specifically whats the size
	fmt.Println(predefined)

	predefined2 := [...]int{1, 2, 3, 4, 45, 5, 6, 77} // ... infers the size from the initialiser
	fmt.Println(predefined2)

	// iterating through array
	for i, v := range predefined2 {
		fmt.Println(i, v)
	}
}
