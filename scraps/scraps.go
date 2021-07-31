package scraps

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Data struct {
	Dolar string
	Euro  string
	Ibov  string
	Cdi   string
	Selic string
	// Balança Comercial
	// PIB
	// PME
	Igpm string
	Ipca string
	Incc string
	Inpc string
}

func leSitesdoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}

func Scraping() Data {
	d := Data{}

	site := leSitesdoArquivo()

	for i, k := range site {

		data, err := http.Get(k)
		if err != nil {
			log.Fatal(err)
		}
		defer data.Body.Close()

		doc, err := goquery.NewDocumentFromReader(data.Body)
		if err != nil {
			log.Fatal(err)
		}

		switch i {

		case 0: //Dolar
			doc.Find("input").Each(func(i int, s *goquery.Selection) {
				if name, _ := s.Attr("name"); name == "currency2" {
					dolar, _ := s.Attr("value")
					d.Dolar = dolar
				}
			})

		case 1: //Euro
			doc.Find("input").Each(func(i int, s *goquery.Selection) {
				if name, _ := s.Attr("name"); name == "currency2" {
					euro, _ := s.Attr("value")
					d.Euro = euro
				}
			})

		case 2: //Ibov
			doc.Find("span").Each(func(i int, s *goquery.Selection) {
				if id, _ := s.Attr("id"); id == "quoteElementPiece1" {
					ibov := s.Text()
					d.Ibov = ibov
				}
			})

		case 3: //CDI
			doc.Find("input").Each(func(i int, s *goquery.Selection) {
				if id, _ := s.Attr("id"); id == "indice-acum" {
					cdi, _ := s.Attr("value")
					d.Cdi = cdi
				}
			})

		case 4: //Selic
			doc.Find("input").Each(func(i int, s *goquery.Selection) {
				if id, _ := s.Attr("id"); id == "selic-ano-meta" {
					selic, _ := s.Attr("value")
					d.Selic = selic
				}
			})

		case 5: //Igpm
			doc.Find("input").Each(func(i int, s *goquery.Selection) {
				if id, _ := s.Attr("id"); id == "indice-acum" {
					igpm, _ := s.Attr("value")
					d.Igpm = igpm
				}
			})

		case 6: //Igpm
			doc.Find("input").Each(func(i int, s *goquery.Selection) {
				if id, _ := s.Attr("id"); id == "indice-acum" {
					ipca, _ := s.Attr("value")
					d.Ipca = ipca
				}
			})

		case 7: //Inpc
			doc.Find("input").Each(func(i int, s *goquery.Selection) {
				if id, _ := s.Attr("id"); id == "indice-acum" {
					inpc, _ := s.Attr("value")
					d.Inpc = inpc
				}
			})

		case 8: //Incc
			doc.Find("input").Each(func(i int, s *goquery.Selection) {
				if id, _ := s.Attr("id"); id == "indice-acum" {
					incc, _ := s.Attr("value")
					d.Incc = incc
				}
			})

		}
	}
	return d
}
