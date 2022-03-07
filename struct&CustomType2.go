package main

import "fmt"

func main() {
	mybill := newBill("willi's bill")

	mybill.addItem("kek", 5.0)
	mybill.addItem("lol", 5.0)
	mybill.addItem("uhh", 5.0)
	mybill.updateTip(10)

	fmt.Println(mybill.format())
}
