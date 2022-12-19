package scraper

import (
	"time"

	"github.com/gocolly/colly"
)

var DollarPrice string

func ScraperDollar() {
	c := colly.NewCollector()

	c.OnHTML("div[class=kf1m0]", func(h *colly.HTMLElement) {

		DollarPrice = h.Text

		time.Sleep(time.Second * 1)
	})

	c.Visit("https://www.google.com/finance/quote/USD-BRL?sa=X&ved=2ahUKEwjznePRlYD8AhXfpZUCHQWqADgQmY0JegQIBhAc")
}
