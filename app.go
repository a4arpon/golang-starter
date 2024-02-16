package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var print = fmt.Println

func getInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, error := reader.ReadString('\n')

	return strings.TrimSpace(input), error
}

func promptOptions(bill Bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose Option \n p = add product \n s = save the bill \n d = add donations for palestine \n Selection : ", reader)

	switch opt {
	case "p":
		print("Please add your product")
		name, _ := getInput("Your Product Name : ", reader)
		price, _ := getInput("Your Product Price : ($) ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			print("Price must be a number.")
			promptOptions(bill)
		}
		bill.addItem(name, p)
		print("Product Added", name, p)
		promptOptions(bill)
	case "s":
		print("You choose to save the bill")
	case "d":
		donation, _ := getInput("Donation For Palestine", reader)
		print("You choose to make a donation", donation)
	default:
		print("Not a valid option")
		promptOptions(bill)
	}
}

func createBill() Bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Customer Name : ", reader)

	bill := newBill(name)

	return bill
}

func main() {
	bill := createBill()
	promptOptions(bill)
	print(bill)
}
