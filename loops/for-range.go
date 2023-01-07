package loops

import (
	"fmt"
)

func PrintArray(tasks []string) {
	for index, val := range tasks {

		fmt.Println(index, val)

		// _, err := os.Open("./test.txt")
		// fmt.Println(err)

	}
}
