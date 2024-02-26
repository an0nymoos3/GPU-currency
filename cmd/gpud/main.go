package main

import (
	"fmt"

	"github.com/an0nymoos3/gpud/internals/gpupricing"
)

func getInput() int {
	var usd int

	fmt.Print("Please enter price in USD: ")
	fmt.Scan(&usd)
	return usd
}

func main() {
	usd_price := getInput()
	gpu_price := gpupricing.GetPrice()
	gpud := float32(usd_price) / float32(gpu_price)

	fmt.Printf("$%d USD = $%f GPUD \n", usd_price, gpud)
}
