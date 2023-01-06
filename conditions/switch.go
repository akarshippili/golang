package conditions

func GetFruitColor(fruit string) (color string) {

	// in go there is a default break after each case
	// to remove use "fallthrough" keyword

	switch fruit {
	case "apple":
		return "red"
	case "orange":
		return "orange"
	case "mango":
		return "yellow"
	case "blueberry":
		return "purple"
	default:
		return ""
	}

}
