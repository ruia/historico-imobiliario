package scrapers

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type Imovirtual struct{}

func (i Imovirtual) ObterPrecosAtualizados(url string) {
	//var preco string
	c := colly.NewCollector(
	//colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "*/*")
	})
	c.OnHTML("strong", func(e *colly.HTMLElement) {
		if e.Attr("aria-label") == "Preço" {
			log.Printf("[Imovirtual] %.0f", ConvertePreco(e.Text))
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "Error:", err)
	})

	c.Visit(url)
	//return preco
}
