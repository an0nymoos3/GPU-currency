package main

import (
	"fmt"

	"github.com/an0nymoos3/gpud/internals/gpupricing"
)

func getInput() float64 {
	var usd float64

	fmt.Print("Please enter price in USD: ")
	fmt.Scan(&usd)
	return usd
}

func main() {
	usd_price := getInput()
	gpu_price := gpupricing.GetPrice()
	gpud := usd_price / float64(gpu_price)

	fmt.Printf("$%f USD = $%f GPUD \n", usd_price, gpud)
}
