package main

import (
	"database/sql"
	// "encoding/json"
	"io"
	"log"

	//"fmt"
	"os"

	"time"

	_ "modernc.org/sqlite"

	"github.com/ruia/historico-imobiliario/scrapers"
)

type AnuncioCompleto struct {
	id               int
	nome             string
	url              string
	imobiliaria_nome string
}

type HistoricoAnuncio struct {
	id         int
	id_anuncio int
	preco      float64
	updated_at string
}

func main() {
	var anuncios []AnuncioCompleto
	var anuncio AnuncioCompleto
	//var idealista scrapers.Idealista
	var imovirtual scrapers.Imovirtual
	//var remax scrapers.Remax

	var historico HistoricoAnuncio
	//var tmp []HistoricoAnuncio

	var preco float64 = 0

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

	for _, anuncio := range anuncios {
		log.Println(anuncio)
		historico.id = 0
		historico.id_anuncio = anuncio.id
		switch anuncio.imobiliaria_nome {
		case "Idealista":
			//preco = idealista.ObterPrecosAtualizados(anuncio.url)
		case "Imovirtual":
			preco = imovirtual.ObterPrecosAtualizados(anuncio.url)
		case "Remax":
			//preco = remax.ObterPrecosAtualizados(anuncio.url)
		}
		historico.preco = preco
		historico.updated_at = time.Now().Format("2006-01-02 15:04:05")
		//temporario
		if preco > 0 {
			//_, err := db.Exec("INSERT INTO historico (id_anuncio, preco, updated_at) VALUES (?, ?, ?);", historico.id_anuncio, historico.preco, historico.updated_at)
			stmt, err := db.Prepare("INSERT INTO historico (id_anuncio, preco, updated_at) VALUES (?, ?, ?)")
			stmt.Exec(historico.id_anuncio, historico.preco, historico.updated_at)
			checkErr(err)
			preco = 0
		}

		//tmp = append(tmp, historico)
	}
	//log.Println(tmp)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("Erro SQLite: %v", err)
		panic(err)
	}
}
