package spec2

import (
	"fmt"
	"math"
)

type MyInt int

func (mi MyInt) double() MyInt {
	return mi * 2
}

func EncapsulationDemo() {
	x := MyInt(343)
	doubled := x.double()
	fmt.Println(doubled)

	p1 := Point{0, 4} // p1 := Point{x: 0, y: 4}
	fmt.Println(p1)
	fmt.Println(p1.distToOrigin())
	p1.offsetX(10)
	fmt.Println(p1)
	fmt.Println(p1.distToOrigin())
}

type Point struct {
	x float64
	y float64
}

func (p Point) distToOrigin() float64 {
	t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
	return math.Sqrt(t)
}

func (p *Point) offsetX(v float64) {
	p.x += v
}
