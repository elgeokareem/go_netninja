package main

import (
	"fmt"
	"math"
	"strings"
)

func sayGreeting(n string) {
	fmt.Printf("Hello %s \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %s \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

// RETURN MULTIPLE VALUES --------------------
func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string

	for _, v := range names {
		initials = append(initials, string(v[0]))
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], ""
}

func main() {
	// sayGreeting("Will")
	// sayBye("Luigi")

	cycleNames([]string{"Mario", "Luigi", "Peach"}, sayGreeting)
	cycleNames([]string{"Mario", "Luigi", "Peach"}, sayBye)

	a1 := circleArea(5)
	a2 := circleArea(5.5)

	fmt.Println(a1, a2)
	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f", a1, a2)

	getInitials("Will vegas")

	// RETURN MULTIPLE VALUES --------------------
	fn1, sn1 := getInitials("Will vegas")
	fmt.Println(fn1, sn1)

	fmt.Println("----------------------")

	SayHello("Will")
}
