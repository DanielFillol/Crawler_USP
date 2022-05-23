package Crawler

import (
	"errors"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"strconv"
	"strings"
)

func resultsFound(doc *html.Node, xpathResult string) (int, error) {
	node, err := htmlquery.Query(doc, xpathResult)
	if err != nil || node == nil {
		return 0, errors.New("xpathResult not found")
	}

	innerText := htmlquery.InnerText(node)
	stringResult1 := strings.Split(innerText, "Resultado: Exibindo")
	stringResult2 := strings.Split(stringResult1[1], " na")
	stringResult3 := strings.Split(stringResult2[0], "de ")
	stringResult4 := stringResult3[1]

	papers, err := strconv.Atoi(stringResult4)
	if err != nil {
		return 0, err
	}

	return papers, nil
}
