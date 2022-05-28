package Crawler

import (
	"errors"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"strconv"
)

func getRow(doc *html.Node, rowNumber int) (DocRow, error) {
	//*[@id="CorpoAcervo"]/div[4]
	node, err := htmlquery.Query(doc, "//*[@id=\"CorpoAcervo\"]/div["+strconv.Itoa(rowNumber)+"]")
	if err != nil {
		return DocRow{}, err
	}

	totalDivs := htmlquery.Find(node, "div")

	//*[@id="CorpoAcervo"]/div[8]
	///div[1]/a

	if len(totalDivs) > 1 {
		name := htmlquery.InnerText(htmlquery.FindOne(node, "div[1]/a"))
		title := htmlquery.InnerText(htmlquery.FindOne(node, "div[2]"))
		field := htmlquery.InnerText(htmlquery.FindOne(node, "div[3]/a"))
		paperType := htmlquery.InnerText(htmlquery.FindOne(node, "div[4]/a"))
		unit := htmlquery.InnerText(htmlquery.FindOne(node, "div[5]/a"))
		year := htmlquery.InnerText(htmlquery.FindOne(node, "div[6]/a"))
		link := htmlquery.SelectAttr(htmlquery.FindOne(node, "div[1]/a"), "href")

		return DocRow{
			Name:  name,
			Title: title,
			Field: field,
			Type:  paperType,
			Unit:  unit,
			Year:  year,
			Link:  link,
		}, nil
	}

	return DocRow{}, errors.New("row not found")
}
