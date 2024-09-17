package spec2

import (
	"fmt"
	"math"
)

func PolymorphismDemo() {
	var shape1 Shape2D
	shape1 = Circle{radius: 7}

	fmt.Println(shape1.area())
	fmt.Println(shape1.perimeter())
	fmt.Println(shape1)

	var rect Rectangle = Rectangle{}
	fmt.Println(FitsInYard(rect))

	var tri Triangle = Triangle{}
	fmt.Println(FitsInYard(tri))

	PrintMe(2.22)
	PrintMe(3)
	PrintMe("hello world")
	PrintMe(tri)
	PrintMe(rect)
	PrintMe(shape1)

	drawShape(rect)
	drawShape(tri)

	shapeIntroduction(rect)
	shapeIntroduction(tri)
}

type Shape2D interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Triangle struct{}

func (t Triangle) area() float64 {
	return 1.1
}

func (t Triangle) perimeter() float64 {
	return 1.1
}

func (t Triangle) introduceMe() {
	fmt.Println("i am a triangle")
}

type Rectangle struct{}

func (r Rectangle) area() float64 {
	return 122.1
}

func (r Rectangle) perimeter() float64 {
	return 122.1
}

func (r Rectangle) whoAmI() {
	fmt.Println("i am a rectangle")
}

func FitsInYard(s Shape2D) bool {
	if s.area() > 100 && s.perimeter() > 100 {
		return true
	}

	return false
}

func PrintMe(val interface{}) { // empty interface allows/accepts all the types
	fmt.Print("value is: ")
	fmt.Println(val)
}

func drawShape(s Shape2D) bool {
	defer fmt.Println("i was called")

	rect, ok := s.(Rectangle) // type assertion [converting from interface to concrete Rectangle type]
	if ok {
		rect.whoAmI()
		return true
	}

	tri, ok := s.(Triangle)
	if ok {
		tri.introduceMe()
		return true
	}

	return false
}

// type switch
func shapeIntroduction(s Shape2D) {
	switch sh := s.(type) {
	case Rectangle:
		sh.whoAmI()
	case Triangle:
		sh.introduceMe()
	default:
		fmt.Println("neither rectangle, nor triangle")
	}
}

// error handling
