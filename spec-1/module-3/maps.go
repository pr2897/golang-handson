package module3

import "fmt"

func MapsDemo() {
	mp := make(map[string]int)
	mp["tr"] = 0
	vl, exists := mp["k"]
	fmt.Println(vl, exists)
	vl, exists = mp["tr"]
	fmt.Println(vl, exists)

}
