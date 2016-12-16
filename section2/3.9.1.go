package main

import (
"fmt"
"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

func main(){
	r1 := Rectangle{12, 2}
	c1 := Circle{25}

	fmt.Printf("Area of r1 is:%f\n", r1.area())
	fmt.Printf("Area of c1 is:%f\n", c1.area())

}
