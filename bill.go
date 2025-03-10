package main

import (
	"fmt"
	"os"
)

type bill struct {
	name string
	items map[string]float64 
	tip float64 
}

func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{},
		tip: 0,
	}

	return b
}

func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	for k, v := range b.items{
		fs += fmt.Sprintf("%25v ...$%v \n", k+":", v)
		total += v 
	}

	fs += fmt.Sprintf("%-25v ...$%v \n", "tip:", b.tip)

	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	return fs
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip  
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price 
}

func (b *bill) save() {
    // Create bills directory if it doesn't exist
    err := os.MkdirAll("bills", 0755)
    if err != nil {
        panic(err)
    }

    data := []byte(b.format())
    err = os.WriteFile("bills/"+b.name+".txt", data, 0644)
    if err != nil {
        panic(err)
    }
    fmt.Println("Bill was saved to bills/" + b.name + ".txt")
}
