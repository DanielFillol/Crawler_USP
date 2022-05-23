package Crawler

import (
	"github.com/antchfx/htmlquery"
	"github.com/tebeka/selenium"
	"golang.org/x/net/html"
	"strconv"
	"strings"
)

func nextPage(driver selenium.WebDriver, followupLink string, i int) (*html.Node, error) {
	link := followupLink + strconv.Itoa(i+1)

	err := driver.Get(link)
	if err != nil {
		return nil, err
	}

	pageSourceCode, err := driver.PageSource()
	if err != nil {
		return nil, err
	}

	document, err := htmlquery.Parse(strings.NewReader(pageSourceCode))
	if err != nil {
		return nil, err
	}

	return document, nil
}
