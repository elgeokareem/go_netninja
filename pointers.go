package main

import "fmt"

func changeName(n *string) {
	fmt.Println("from changeName", *n)
	*n = "lulz"
	fmt.Println("from changeName", *n)
}

func main() {
	name := "kekerinos"
	namePointer := &name

	changeName(namePointer)

	fmt.Println("from main", name)
}
