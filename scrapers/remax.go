package scrapers

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type Remax struct{}

func (i Remax) ObterPrecosAtualizados(url string) {
	//var preço string
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.82 Safari/537.36"),
	)
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "*/*")
	})
	c.OnHTML("input", func(e *colly.HTMLElement) {
		// log.Printf("[Remax] %.0f", ConvertePreco(e.ChildText(".listing-price")))
		log.Printf("[Remax]", e)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "Error:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.Visit(url)
	//return preço
}
