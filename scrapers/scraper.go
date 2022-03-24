package scrapers

type Scraper interface {
	ObterPrecosAtualizados(url string)
}
