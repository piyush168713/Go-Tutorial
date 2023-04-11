package main
import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rect struct {
	height float64
	width float64
}

func (r rect) area() float64 {
	return r.height * r.width;
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius;
}

func getArea(s shape) float64 {
	return s.area();
}

func main(){
	/*
	c1 := circle{4.1}
	r1 := rect{3,4}
 
	shapes := []shape{c1,r1}

	for _, shape2 := range shapes {
		// fmt.Println(shape2.area())
		fmt.Println(getArea(shape2))
	}
	*/
	c1 := circle{4.1}
	r1 := rect{3,4}

	shapes := []shape{c1,r1}

	for _, sh2 := range shapes {
		fmt.Println(getArea(sh2))
	}

}