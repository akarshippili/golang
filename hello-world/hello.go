package main

import (
	"fmt"

	"github.com/akarshippili/golang/conditions"
	"github.com/akarshippili/golang/functions"
	"github.com/akarshippili/golang/variables"
)

func main() {
	fmt.Println("Hello World!!!")

	variables.Test()
	fmt.Println(variables.TestVariable)

	x := "akarsh ippili"
	y := "how to become a top-g (lol!)"
	author, course := functions.Converter(x, y)

	fmt.Println(x, y)
	fmt.Println(author, course)

	fmt.Println(functions.Sum(1, 2, 3))
	fmt.Println(functions.Sum(-3, -2, -1, 0, 1, 2, 3))

	conditions.PrintEvenOdd(1)
	conditions.PrintEvenOdd(2)

	fmt.Println("apple's color:", conditions.GetFruitColor("apple"))
	fmt.Println("mango's color:", conditions.GetFruitColor("mango"))
}
