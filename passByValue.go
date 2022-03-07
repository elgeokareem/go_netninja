package main

import "fmt"

func updateName(x string) string {
	x = "New Name"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.5
}

func main() {
	// Group A: strins, ints, bools, floats, arrays structs
	// Non-pointer-values
	name := "Old Name"

	name = updateName(name)

	fmt.Println(name)

	// Group B: slices, maps, functions
	// Pointer-wrapper-values
	menu := map[string]float64{
		"pie":  5.0,
		"cake": 10.0,
	}

	updateMenu(menu)

	fmt.Println(menu["coffee"])
}
