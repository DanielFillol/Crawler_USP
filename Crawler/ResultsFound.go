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

	// TODO: not the best way

	replaceText1 := strings.Replace(innerText, "Resultado: Exibindo", "", -1)
	lastChars := replaceText1[len(replaceText1)-18:]
	mainText := strings.Replace(replaceText1, lastChars, "", -1)
	final := strings.Replace(mainText, mainText[0:7], "", -1)
	papersString := strings.TrimSpace(final)
	s := strings.Split(papersString, " ")

	papers, err := strconv.Atoi(s[0])
	if err != nil {
		return 0, err
	}

	return papers, nil
}
