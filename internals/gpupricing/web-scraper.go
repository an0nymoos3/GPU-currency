package gpupricing

import (
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func GetPrice() int {
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
