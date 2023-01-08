package lists

import (
	"fmt"
)

func TestLists() {
	// array - fixed size like arrays in java
	// slices - resizeable like arrayList

	// one way to declare
	// courses := make([]string, 5, 10)
	// fmt.Printf("Lenght of slice is %d.\ncapacity of cources is %d.\n", len(courses), cap(courses))

	// courses[0] = "Go Fundamentels"
	// courses[1] = "AWS by java techie"
	// courses[2] = "Networking Fundamentals"
	// courses[3] = "Go Web Applications, Go net, syscall packages"
	// courses[4] = "Kubernates or K8s"

	// cleaner way to declare
	// when declared in this way len and cap are equal to the no. of elements.
	courses := []string{
		"Go Fundamentels",
		"AWS by java techie",
		"Networking Fundamentals",
		"Go Web Applications, Go net, syscall packages",
		"Kubernates or K8s by nana",
	}
	fmt.Printf("Addr: %p.\nLenght of slice is %d.\ncapacity of cources is %d.\n", &courses, len(courses), cap(courses))

	courses = append(courses, "Kubernates by nigel poulton")
	fmt.Printf("Addr: %p.\nLenght of slice is %d.\ncapacity of cources is %d.\n", &courses, len(courses), cap(courses))

	for index, course := range courses {
		fmt.Printf("%d. %s.\n", index, course)
	}

	// slice
	fmt.Println(courses[1:3])
}
