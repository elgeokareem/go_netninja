package main

import "fmt"

func main() {
	menu := map[string]float64{
		"Hamburger": 4.00,
		"Fries":     2.00,
		"Drink":     1.50,
	}

	fmt.Println(menu)
	fmt.Println(menu["Hamburger"])

	// looping maps
	for key, value := range menu {
		fmt.Println(key, value)
	}
}
