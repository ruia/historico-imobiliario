package scrapers

import (
	// "fmt"
	"github.com/gocolly/colly"
	"log"
)

type Idealista struct {}

func (i Idealista) ObterPrecosAtualizados(url string) {
	//var preco string
	c := colly.NewCollector()

	c.OnHTML("span[class=info-data-price]>span[class=txt-bold]", func(e *colly.HTMLElement) {
		// log.Printf("[Auchan] %.2f € \n", convert(e.ChildText(".value")))
		log.Printf("[Imovirtual]", e.ChildText()))
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "Error:", err)
	})

	c.Visit(url)
	//return preco
}