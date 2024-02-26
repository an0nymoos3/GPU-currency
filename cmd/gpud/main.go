package main

import (
	"fmt"

	"github.com/an0nymoos3/gpud/internals/gpupricing"
)

func getInput() (float64, string) {
	var price float64
	var currency string

	// fmt.Scan() scans space seperated tokens, which works well for price and currency input
	fmt.Print("Please enter the price in the format price currency (eg. 500 USD): ")
	fmt.Scan(&price)
	fmt.Scan(&currency)

	// Return price as f64 and currency as string
	return price, currency
}

func main() {
	price, currency := getInput()
	gpud := gpupricing.ConvertToGpud(price, currency)

	fmt.Printf("$%f %s = $%f GPUD \n", price, currency, gpud)
}
