package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add Item, s - save bill, t- add tip): ", reader)

	switch opt {
	case "a":
		fmt.Println("Add item to the bill")
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("Invalid price")
			promptOptions(b)
			return
		}

		b.addItem(name, p)

		fmt.Println("Item added - ", name, price)
		promptOptions(b)

	case "s":
		b.save()
		fmt.Println("Saved bill", b)

	case "t":
		fmt.Println("Add tip")
		tip, _ := getInput("Enter Tip amount: ", reader)

		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("Invalid tip")
			promptOptions(b)
			return
		}

		b.updateTip(t)

		fmt.Println("Tip added - ", tip)
		promptOptions(b)

	default:
		fmt.Println("Invalid option...")
		promptOptions(b)
	}
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func main() {
	mybill := createBill()

	promptOptions(mybill)
	fmt.Println(mybill)
}
