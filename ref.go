package main

import "fmt"

type Bill struct {
	name     string
	items    map[string]float64
	donation float64
}

func newBill(name string) Bill {
	bill := Bill{
		name:     name,
		items:    map[string]float64{},
		donation: 0,
	}
	return bill
}

func (bill *Bill) format() string {
	fs := "Bill Breakdown \n"
	var total float64 = 0
	for k, p := range bill.items {
		fs += fmt.Sprintf("%-25v $%v \n", k+":", p)
		total += p
	}

	fs += fmt.Sprintf("%-25v $%v\n", "Donation: ", bill.donation)

	fs += fmt.Sprintf("%-25v $%0.2f", "Total: ", total+bill.donation)
	return fs
}

func (bill *Bill) charityForPalestine(donation float64) {
	bill.donation = donation
}

func (bill *Bill) addItem(name string, price float64) {
	bill.items[name] = price
}
