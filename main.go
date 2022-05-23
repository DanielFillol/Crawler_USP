package main

import (
	"fmt"
	"github.com/Darklabel91/Crawler_USP/CSV"
	"github.com/Darklabel91/Crawler_USP/Crawler"
	"github.com/tebeka/selenium"
)

const (
	FUP          = "&pagina="
	XpathMasters = "//*[@id=\"CorpoAcervo\"]/div[3]"
	XpathPapers  = "//*[@id=\"CorpoAcervo\"]/div[1]"
	MastersInit  = 6
	PapersInit   = 4
)

var crawlerList = []string{
	"http://www.tcc.sc.usp.br/index.php?option=com_jumi&fileid=17&Itemid=178&lang=br&id=89&curso=89001&hab=0",
	"https://www.teses.usp.br/index.php?option=com_jumi&fileid=9&Itemid=159&lang=pt-br&id=107131&prog=107001&exp=0",
	"https://www.teses.usp.br/index.php?option=com_jumi&fileid=9&Itemid=159&lang=pt-br&id=2131&prog=2001&exp=0",
	"https://www.teses.usp.br/index.php?option=com_jumi&fileid=9&Itemid=159&lang=pt-br&id=2132&prog=2001&exp=0",
	"https://www.teses.usp.br/index.php?option=com_jumi&fileid=9&Itemid=159&lang=pt-br&id=2134&prog=2001&exp=0",
	"https://www.teses.usp.br/index.php?option=com_jumi&fileid=9&Itemid=159&lang=pt-br&id=2138&prog=2001&exp=0",
	"https://www.teses.usp.br/index.php?option=com_jumi&fileid=9&Itemid=159&lang=pt-br&id=2138&prog=2001&exp=0",
	"https://www.teses.usp.br/index.php?option=com_jumi&fileid=9&Itemid=159&lang=pt-br&id=2133&prog=2001&exp=0",
	"https://www.teses.usp.br/index.php?option=com_jumi&fileid=9&Itemid=159&lang=pt-br&id=2135&prog=2001&exp=0",
	"https://www.teses.usp.br/index.php?option=com_jumi&fileid=9&Itemid=159&lang=pt-br&id=2136&prog=2001&exp=0",
	"https://www.teses.usp.br/index.php?option=com_jumi&fileid=9&Itemid=159&lang=pt-br&id=107&prog=1&exp=0",
}

func main() {
	driver, err := Crawler.WebDriver()
	if err != nil {
		fmt.Println(err)
	}
	defer driver.Close()

	data, err := crawProcedure(driver, crawlerList)
	if err != nil {
		fmt.Println(err)
	}

	err = CSV.WriteCSV("USP", "Result", data)
	if err != nil {
		fmt.Println(err)
	}
}

func crawProcedure(driver selenium.WebDriver, crawlerList []string) ([]Crawler.DocRow, error) {
	var data []Crawler.DocRow
	for i := 0; i < len(crawlerList); i++ {
		fup := crawlerList[i] + FUP

		if i == 0 {
			pprs, err := Crawler.Craw(driver, crawlerList[i], fup, XpathPapers, PapersInit)
			if err != nil {
				return nil, err
			}

			for _, ppr := range pprs {
				data = append(data, ppr)
			}

		} else {
			mstrs, err := Crawler.Craw(driver, crawlerList[i], fup, XpathMasters, MastersInit)
			if err != nil {
				return nil, err
			}

			for _, mstr := range mstrs {
				data = append(data, mstr)
			}
		}
	}

	return data, nil
}
