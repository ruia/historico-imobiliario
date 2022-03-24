package main

import (
	"database/sql"
	// "encoding/json"
	"io"
	"log"

	//"fmt"
	"os"

	_ "modernc.org/sqlite"

	"github.com/ruia/historico-imobiliario/scrapers"
)

type AnuncioCompleto struct {
	id               int
	nome             string
	url              string
	imobiliaria_nome string
}

func main() {
	var anuncios []AnuncioCompleto
	var anuncio AnuncioCompleto

	f, err := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	log.Println("App started")

	db, err := sql.Open("sqlite", "./dados.db")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM anuncios_completo")
	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&anuncio.id, &anuncio.nome, &anuncio.url, &anuncio.imobiliaria_nome)
		checkErr(err)
		anuncios = append(anuncios, anuncio)
	}
	log.Println("Foram carregados", len(anuncios), "anuncio(s)")
	//var url string
	var idealista scrapers.Idealista
	var imovirtual scrapers.Imovirtual
	var remax scrapers.Remax
	//scrappers := []scrapers.Scraper{scrapers.Idealista{}}

	for _, anuncio := range anuncios {
		log.Println(anuncio)
		switch anuncio.imobiliaria_nome {
		case "Idealista":
			idealista.ObterPrecosAtualizados(anuncio.url)
		case "Imovirtual":
			imovirtual.ObterPrecosAtualizados(anuncio.url)
		case "Remax":
			remax.ObterPrecosAtualizados(anuncio.url)
		}
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("Erro SQLite: %v", err)
		panic(err)
	}
}
