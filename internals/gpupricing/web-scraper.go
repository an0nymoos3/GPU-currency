package gpupricing

import (
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func GetGPUPrice() int {
	int_price := 0
	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
		panic(err)
	})

	c.OnHTML("div.mt-2.font-bold.text-4xl", func(e *colly.HTMLElement) {
		price := string(e.Text)
		price = strings.ReplaceAll(price, "$", "")

		// TODO: Replace with prettier solution to avoid overwriting retail price with 2nd hand price.
		if int_price == 0 { // Avoid overwriting retail price with second hand price
			var err error
			int_price, err = strconv.Atoi(price)
			if err != nil { // If error
				panic(err)
			}
		}
	})

	c.Visit("https://bestvaluegpu.com/history/new-and-used-rtx-4080-price-history-and-specs/")

	return int_price
}

func XeConvertToUsd(price float64, currency string) float64 {
	usd_price := 0.0
	link := generateConversionLink(price, currency)
	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Failed to reach: ", link)
		panic(err)
	})

	c.OnHTML("p.result__BigRate-sc-1bsijpp-1.dPdXSB", func(e *colly.HTMLElement) {
		substrings := strings.Split(e.Text, " ")

		var err error
		usd_price, err = strconv.ParseFloat(substrings[0], 64)
		if err != nil {
			panic(err)
		}
	})

	c.Visit(link)

	return usd_price
}

func generateConversionLink(price float64, currency string) string {
	link := "https://www.xe.com/currencyconverter/convert/?Amount="
	link += strconv.FormatFloat(price, 'f', 3, 64)
	link += "&From=" + string(currency)
	link += "&To=USD"

	return link
}
