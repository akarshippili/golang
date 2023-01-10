package main

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/akarshippili/golang/concurrency"
	"github.com/akarshippili/golang/conditions"
	ds "github.com/akarshippili/golang/datastructures"
	"github.com/akarshippili/golang/functions"
	"github.com/akarshippili/golang/lists"
	"github.com/akarshippili/golang/loops"
	"github.com/akarshippili/golang/maps"
	"github.com/akarshippili/golang/models"
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

	conditions.TryRead()

	lists.TestLists()

	loops.CountDown(10)
	loops.PrintArray([]string{"learn go", "networking tcp", "revise design patterns", "leetcode daliy question and contest"})

	// goCourse := new(structs.Course)
	goCourse := models.Course{
		Author: "akarsh",
		Level:  "beginner",
		Rating: 4.26,
	}

	fmt.Println("type: ", reflect.TypeOf(goCourse), "course: ", goCourse)

	loops.CountDown(3)
	loops.PrintArray([]string{"learn go", "networking tcp", "revise design patterns", "leetcode daliy question and contest"})

	maps.TestMaps()

	ds.BtreeTest()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		concurrency.One()
		wg.Done()
	}()

	go func() {
		concurrency.Two()
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("main goroutine exiting")
}
