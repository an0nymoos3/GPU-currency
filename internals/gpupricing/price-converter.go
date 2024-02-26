package gpupricing

func ConvertToGpud(price float64, currency string) float64 {
	if currency != "USD" {
		price = XeConvertToUsd(price, currency)
	}
	return price / float64(GetGPUPrice())
}
