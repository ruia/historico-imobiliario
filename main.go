package main

import (
    "database/sql"
    // "encoding/json"
     "io"
    "log"
    //"fmt"
    "os"
    _ "modernc.org/sqlite"
)

type AnuncioCompleto struct {
    id                  int
    nome                string
    url                 string
    imobiliaria_nome    string
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

    for _, anuncio := range anuncios {
        log.Println(anuncio)

    }


    //fmt.Println(anuncios)
    
    // products := getProductsList()

    // scrappers := []scrapers.Scraper{scrapers.Worten{}, scrapers.RadioPopular{}, scrapers.MediaMarket{}, scrapers.Auchan{}, scrapers.CasaCarvalho{}, scrapers.CastroElectronica{}, scrapers.Prinfor{}}

    // for _, product := range products.Products {
    // 	log.Println(product.Name)
    // 	log.Println("-----------------")
    // 	for _, scrapper := range scrappers {
    // 		process(scrapper, product)
    // 	}
    // 	log.Println("")
    // }

}

func checkErr(err error) {
    if err != nil {
        log.Fatalf("Erro SQLite: %v", err)
        panic(err)
    }
}