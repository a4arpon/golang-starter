package main

import (
	"fmt"
	"os"
)

type Bill struct {
	name     string
	phone    string
	items    map[string]float64
	donation float64
}

func newBill(name string, phone string) Bill {
	bill := Bill{
		name:     name,
		phone:    phone,
		items:    map[string]float64{},
		donation: 0,
	}
	return bill
}

func (bill *Bill) format() string {
	fs := "Customer : " + bill.name + "\n" + "Phone : " + bill.phone + "\n" + "Bill Breakdown \n"
	var total float64 = 0
	for k, p := range bill.items {
		fs += fmt.Sprintf("%-32v $%v \n", k+":", p)
		total += p
	}

	fs += fmt.Sprintf("%-32v $%v\n", "Donation: ", bill.donation)

	fs += fmt.Sprintf("%-32v $%0.2f", "Total: ", total+bill.donation)
	return fs
}

func (bill *Bill) charityForPalestine(donation float64) {
	bill.donation += donation
}

func (bill *Bill) addItem(name string, price float64) {
	bill.items[name] = price
}

func (bill *Bill) saveBill() {
	data := []byte(bill.format())
	err := os.WriteFile("bills/"+bill.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("Bill was saved to file.")
}
