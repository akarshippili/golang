package variables

import "fmt"

func test() {
	name := "test"

	fmt.Println("int test")
	fmt.Println(name, &name)

	test2(name)
}

func test2(name string) {
	fmt.Println("in test2")
	fmt.Println(name, &name)
}
