package functions

import "strings"

func Converter(author string, course string) (a string, c string) {
	author = strings.ToUpper(author)
	course = strings.ToUpper(course)

	return author, course
}
