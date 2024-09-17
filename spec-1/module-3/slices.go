package module3

import "fmt"

func SliceDemo() {
	arr := [...]string{"a", "b", "c", "d", "e", "f"}
	s1 := arr[1:3]
	s2 := arr[2:4]

	s1[1] = "piyush" // slice refers to same underlying array. changing it can affects other slices.

	fmt.Println(arr)
	fmt.Println(s1)
	fmt.Println(s2)

	slic := make([]int, 10)      // type, capacity
	slice := make([]int, 10, 15) // type, size, capacity

	fmt.Println(slic, slice)

	slice = append(slice, 11)

	fmt.Println(slic, slice)

}
