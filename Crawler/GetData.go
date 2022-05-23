package Crawler

import (
	"github.com/antchfx/htmlquery"
	"github.com/tebeka/selenium"
	"golang.org/x/net/html"
	"strings"
)

type DocRow struct {
	Name  string
	Title string
	Field string
	Type  string
	Unit  string
	Year  string
	Link  string
}

func GetData(driver selenium.WebDriver, document *html.Node, pageAmount int, followupLink string, initTable int) ([]DocRow, error) {
	var docs []DocRow
	var err error
	for i := 0; i < pageAmount; i++ {

		if i != 0 {
			document, err = nextPage(driver, followupLink, i)
			if err != nil {
				return nil, err
			}
		}

		divs := htmlquery.Find(document, "//*[@id=\"CorpoAcervo\"]/div")
		loop := len(divs) - (initTable + 1)

		for j := 0; j < loop; j++ {
			ppr, err := getRow(document, initTable+j)
			if err != nil {
				return nil, err
			}
			if strings.Contains(strings.ToLower(ppr.Field), "direito") {
				docs = append(docs, ppr)
			}
		}
	}

	return docs, nil
}
