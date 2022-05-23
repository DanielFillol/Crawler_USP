package Crawler

import (
	"github.com/antchfx/htmlquery"
	"github.com/tebeka/selenium"
	"strings"
)

const ResultByPage = 10

func Craw(driver selenium.WebDriver, mainLink string, followupLink string, xpathResult string, initTable int) ([]DocRow, error) {
	err := driver.Get(mainLink)
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

	pages, err := resultsFound(document, xpathResult)
	if err != nil {
		return nil, err
	}

	pageAmt := pageAmount(pages)

	data, err := GetData(driver, document, pageAmt, followupLink, initTable)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func pageAmount(pages int) int {
	amount := pages / ResultByPage
	if pages%ResultByPage != 0 {
		amount = amount + 1
	}
	return amount
}
