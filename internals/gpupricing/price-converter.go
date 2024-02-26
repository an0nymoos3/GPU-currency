package gpupricing

import "fmt"

func ConvertToGpud(price float64, currency string) float64 {
	if currency != "USD" {
		fmt.Println("Sending currency conversion request...")
		price = XeConvertToUsd(price, currency)
	}
	return price / float64(GetGPUPrice())
}
