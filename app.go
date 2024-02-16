package main

import (
	"fmt"
)

var print = fmt.Println

func createBill() {

}

func main() {
	myBill := newBill("Lisa")

	myBill.addItem("Note Book", 5)

	myBill.charityForPalestine(2)

	print(myBill.format())
}
