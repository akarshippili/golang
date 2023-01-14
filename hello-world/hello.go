package main

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/akarshippili/golang/concurrency"
	"github.com/akarshippili/golang/conditions"
	"github.com/akarshippili/golang/datastructures"
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

	uf := datastructures.GetUnionFind(10)

	uf.Merge(0, 1)
	uf.Merge(3, 1)

	uf.Merge(2, 4)
	uf.Merge(2, 6)
	uf.Merge(2, 8)
	uf.Merge(1, 1)

	fmt.Println("num of components", uf.GetNumOfComp())
	fmt.Println("are 2 and 1 connected", uf.IsConnected(1, 2))
	fmt.Println("are 2 and 8 connected", uf.IsConnected(8, 2))
}
