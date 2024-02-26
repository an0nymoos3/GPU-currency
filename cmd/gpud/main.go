package main

import (
	"fmt"

	"github.com/an0nymoos3/gpud/internals/gpupricing"
)

func getInput() float32 {
	var usd float32

	fmt.Print("Please enter price in USD: ")
	fmt.Scan(&usd)
	return usd
}

func main() {
	usd_price := getInput()
	gpu_price := gpupricing.GetPrice()
	gpud := usd_price / float32(gpu_price)

	fmt.Printf("$%f USD = $%f GPUD \n", usd_price, gpud)
}
