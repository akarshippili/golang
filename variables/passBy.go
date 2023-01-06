package variables

import "fmt"

func Test() {
	name := "test"

	fmt.Println("int test")
	fmt.Println(name, &name)

	Test2(name)
	Test3(&name)
}

func Test2(name string) {
	fmt.Println("in test2")
	fmt.Println(name, &name)
}

func Test3(ptr *string) {
	fmt.Println("in test3")
	fmt.Println(*ptr, ptr)
}
