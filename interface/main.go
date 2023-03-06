package main

import (
	"fmt"
	"math"
)

type Shape interface {
	perimeter() float64
	area() float64
}

func printShapeDetails(shape Shape) {
	fmt.Println("Shape", shape)
	fmt.Println("Shape", shape.perimeter())
	fmt.Println("Shape", shape.area())
}

type Rectangle struct {
	width  float64
	length float64
}

func (rectangle *Rectangle) perimeter() float64 {
	return 2 * (rectangle.width + rectangle.length)
}

func (rectangle *Rectangle) area() float64 {
	return rectangle.width * rectangle.length
}

type Circle struct {
	radius float64
}

func (circle *Circle) perimeter() float64 {
	return 2 * math.Pi * circle.radius
}

func (circle *Circle) area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}

func main() {
	rectangle := Rectangle{
		width:  10,
		length: 12,
	}
	printShapeDetails(&rectangle)

	circle := Circle{
		radius: 2,
	}
	printShapeDetails(&circle)
}
