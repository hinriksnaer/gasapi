package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func getGasData(gasType string) {
	fName := gasType + ".csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Instantiate default collector
	c := colly.NewCollector()

	c.OnHTML("#okt95 tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr + tr", func(_ int, el *colly.HTMLElement) {
			writer.Write([]string{
				el.ChildText("td:nth-child(1)"),
				el.ChildText("td:nth-child(2)"),
				el.ChildText("td:nth-child(3)"),
			})
		})
	})

	c.Visit("http://gsmbensin.is/gsmbensin_web.php")

	log.Printf("Scraping finished, check file %q for results\n", fName)
}
