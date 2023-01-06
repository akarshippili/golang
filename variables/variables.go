package variables

import (
	"fmt"
	"reflect"
	"strconv"
)

// package level variables
var (
	name   string // initilizes with empty string
	course string = "getting started with go"
	module int    // initilizes with zero
	clip   = 2    // type is infered as int
)

func testing() {
	fmt.Println(name, course, module, clip)
	fmt.Println("name type: ", reflect.TypeOf(name))
	fmt.Println("clip type: ", reflect.TypeOf(clip))

	// var a int = 10
	// var a = 10
	// := short assignment method
	a := 10
	fmt.Println(a)

	val, err := strconv.Atoi("10")

	if err == nil {
		fmt.Println(val)
	} else {
		fmt.Println(err)
	}

	fmt.Println("Memory address of *course* variable is ", &course)

	var ptr *string = &course
	fmt.Println("Poiniting course variable at address,", ptr, "which holds this value,", *ptr)
	fmt.Println("ptr address", &ptr)
}
