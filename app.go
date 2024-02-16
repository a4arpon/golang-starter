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
		bill.saveBill()
		print("You choose to save the bill ")

	case "d":
		donation, _ := getInput("Donation For Palestine ($) ", reader)
		d, err := strconv.ParseFloat(donation, 64)
		if err != nil {
			print("Donation must be a number.")
			promptOptions(bill)
		}
		bill.charityForPalestine(d)
		print("You choose to make a donation", d)
		promptOptions(bill)

	default:
		print("Not a valid option")
		promptOptions(bill)
	}
}

func createBill() Bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Customer Name : ", reader)
	phone, _ := getInput("Customer Phone : ", reader)

	bill := newBill(name, phone)

	return bill
}

func main() {
	bill := createBill()
	promptOptions(bill)
	print(bill)
}
