package maps

import "fmt"

func TestMaps() {
	alphabets := make(map[string]int)

	alphabets["d"] = 2
	alphabets["c"] = 2
	alphabets["e"] = 2
	alphabets["f"] = 2
	alphabets["a"] = 1
	alphabets["b"] = 2

	// prints in a sorted manner
	fmt.Println("alphabets mapping", alphabets)

	// in random order while using range loop
	for key, val := range alphabets {
		fmt.Printf("key : %v, value: %v\n", key, val)
	}

	// accessing
	// map[key]

	// updating
	// map[key] = newValue

	// delete
	//delete(map, key)
}
