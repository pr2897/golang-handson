package spec2

import (
	"fmt"
	"math"
)

var person1 Person

func FunctionDemo() {
	increment := foo(5)
	fmt.Println(increment)

	person1 = Person{
		Name:    "Piyush",
		Address: Address{Road: "MG Road", Pincode: "851204", Building: "HPC, khagaria"},
		Course:  []Course{{Title: "Comp sc.", Duration: "33:33"}},
	}

	original, incremented, incremented2times := foo2(5)
	fmt.Println(original, incremented, incremented2times)

	x := 3
	foo3(x)
	fmt.Println(x)

	DemoCallByValue()
	DemoCallByReference()
	firstClassFunctionDemo()

	demoNamedReturnFunction()

	closureDemo()
}

// single return
func foo(x int) int {
	return x + 1
}

// multi return
func foo2(x int) (int, int, int) {
	return x, x + 1, x + 2
}

// call by value is done is go (arguments that are passed as parameters, they are copied to parameters)
// modifying parameters has no effects outside the function.

func foo3(y int) {
	y = y + 1
	println(y)
}

type Person struct {
	Name    string
	Address Address
	Course  []Course
}

type Address struct {
	Building string
	Road     string
	Pincode  string
}

type Course struct {
	Title    string
	Duration string
}

func callByValue(person Person) Person {
	person.Name = "ayush"
	return person
}

func foo4(x [3]int) [3]int {
	x[1] = 2
	x[0] = 12

	return x
}

func DemoCallByValue() {
	var final Person
	fmt.Println(person1, final)
	final = callByValue(person1)
	fmt.Println(person1, final)

	initialArr := [3]int{1, 2, 3}
	var finalArr [3]int

	fmt.Println(initialArr, finalArr)
	finalArr = foo4(initialArr) // initialArr got copied to parameter and changes happened on parameter.
	fmt.Println(initialArr, finalArr)

}

func DemoCallByReference() {
	// passing array pointer
	// using array
	a := [3]int{1, 2, 3}
	fmt.Println(a)
	foo5(&a)
	fmt.Println(a)

	// using slice
	sl := []int{1, 2, 3} // []int{} this declares slice whereas, [3]int{} declares array of size 3.
	fmt.Println(sl)
	foo6(sl)
	fmt.Println(sl)
}

func foo6(x []int) {
	x[0] = 12
}

func foo5(x *[3]int) int {
	(*x)[0] = (*x)[0] + 1

	return 2
}

func firstClassFunctionDemo() {
	var funcVar func(int) int = func(i int) int {
		return i + 1
	}

	fmt.Println(funcVar(2))

	v := applyIt(func(i int) int { return i * 10 }, 2)
	fmt.Println(v)

	dist1 := makeDistOrigin(0, 0)
	dist2 := makeDistOrigin(2, 2)

	fmt.Println(dist1(2, 2))
	fmt.Println(dist2(2, 2))

	// defer function demo
	deferFunctionDemo()

	newX, newY, newZ := conversions(func(x int) int { return x * 2 }, 1, 2, 3)
	fmt.Println(newX, newY, newZ)
}

// function taking input as function
func applyIt(afunct func(int) int, val int) int {
	return afunct(val)
}

// function returning function
func makeDistOrigin(o_x, o_y float64) func(float64, float64) float64 {
	fn := func(x, y float64) float64 {
		return math.Sqrt(math.Pow(x-o_x, 2) + math.Pow(y-o_y, 2))
	}

	return fn
}

// defer function
func deferFunctionDemo() {
	fmt.Println("Hello world")
	defer fmt.Println("exiting from here")
	defer fmt.Println("cleanup")

	fmt.Println("what am i doing")
	defer fmt.Println("testing")
}

// named return
func demoNamedReturnFunction() {
	fmt.Println("---- named functions here ----")
	fmt.Println(getCords())
	fmt.Println("---- named functions ends here ----")
}

func yearsUntilEvents(age int) (yearsUntilAdult int, yearsUntilDrinking int, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return
}

func getCords() (x, y int) {
	x = 5
	y = 7
	return // naked return
	// return y, x // explict return
}

// anonymous functions
func conversions(converter func(int) int, x, y, z int) (int, int, int) {
	convertedX := converter(x)
	convertedY := converter(y)
	convertedZ := converter(z)
	return convertedX, convertedY, convertedZ
}

// closure
func closureDemo() {
	messageAgg := concatter()
	message := messageAgg("hello")
	fmt.Println(message)
	message = messageAgg("world")
	fmt.Println(message)
}

func concatter() func(string) string {
	doc := ""

	return func(word string) string {
		doc += word + " "
		return doc
	}
}
