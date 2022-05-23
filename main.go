package main

import (
	"fmt"
	"github.com/Darklabel91/Crawler_USP/CSV"
	"github.com/Darklabel91/Crawler_USP/Crawler"
)

const (
	Masters       = "https://www.teses.usp.br/index.php?option=com_jumi&fileid=7&Itemid=62&lang=pt-br"
	MastersFUP    = Masters + "&pagina="
	ResultMasters = "//*[@id=\"CorpoAcervo\"]/div[3]"
	MastersInit   = 6
	Papers        = "http://www.tcc.sc.usp.br/index.php?option=com_jumi&fileid=17&Itemid=178&lang=br&id=89&curso=89001&hab=0"
	PapersFUP     = Papers + "&pagina="
	ResultPapers  = "//*[@id=\"CorpoAcervo\"]/div[1]"
	PapersInit    = 4
)

func main() {
	driver, err := Crawler.WebDriver()
	if err != nil {
		fmt.Println(err)
	}
	defer driver.Close()

	papersLaw, err := Crawler.Craw(driver, Papers, PapersFUP, ResultPapers, PapersInit)
	if err != nil {
		fmt.Println(err)
	}

	masterCivilLaw, err := Crawler.Craw(driver, Masters, MastersFUP, ResultMasters, MastersInit)
	if err != nil {
		fmt.Println(err)
	}

	data := joinData(papersLaw, masterCivilLaw)

	err = CSV.WriteCSV("USP", "Result", data)
	if err != nil {
		fmt.Println(err)
	}
}

func joinData(data1 []Crawler.DocRow, data2 []Crawler.DocRow) []Crawler.DocRow {
	var data []Crawler.DocRow
	for _, dt := range data1 {
		data = append(data, dt)
	}
	for _, dt := range data2 {
		data = append(data, dt)
	}
	return data
}
