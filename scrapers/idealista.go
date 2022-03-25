package scrapers

import (
	// "fmt"
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type Idealista struct{}

func (i Idealista) ObterPrecosAtualizados(url string) float64 {
	var preco float64

	c := colly.NewCollector()

	c.OnHTML("span[class=info-data-price]>span[class=txt-bold]", func(e *colly.HTMLElement) {
		preco = ConvertePreco(e.ChildText(".value"))
		log.Printf("[Imovirtual] %.2f € \n", preco)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "Error:", err)
	})

	c.Visit(url)

	return preco
}
