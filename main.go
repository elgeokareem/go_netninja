package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// var ages [3]int = [3]int{10, 20, 30}
	var ages = [3]int{10, 20, 30}
	names := [3]string{"a", "b", "c"}
	names[1] = "d"

	fmt.Println(ages)
	fmt.Println(names)

	// SLICES
	var scores = []int{10, 20, 30}
	scores[2] = 25
	scores = append(scores, 40)
	fmt.Println(scores)

	// Slices ranges
	corte := scores[:2]
	fmt.Println(corte)

	// Libreria standard strings ------------------
	var str = "Hello World"
	fmt.Println(strings.ReplaceAll(str, "l", "x"))

	// ---------
	agesTwo := [3]int{35, 20, 10}
	sort.Ints(agesTwo[:]) // altera el array original
	fmt.Println(agesTwo)

	index := sort.SearchInts(agesTwo[:], 20)
	fmt.Println(index)
}
