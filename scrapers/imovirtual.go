package scrapers

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type Imovirtual struct{}

func (i Imovirtual) ObterPrecosAtualizados(url string) float64 {
	var preco float64

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "*/*")
	})

	c.OnHTML("strong", func(e *colly.HTMLElement) {
		if e.Attr("aria-label") == "Preço" {
			preco = ConvertePreco(e.Text)
			log.Printf("[Imovirtual] %.0f", preco)
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "Error:", err)
	})

	c.Visit(url)

	return preco
}
