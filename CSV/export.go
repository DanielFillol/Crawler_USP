package CSV

import (
	"encoding/csv"
	"github.com/Darklabel91/Crawler_USP/Crawler"
	"os"
	"path/filepath"
)

func WriteCSV(fileName string, folderName string, papers []Crawler.DocRow) error {
	var rows [][]string

	rows = append(rows, generateHeaders())

	for _, decision := range papers {
		rows = append(rows, generateRow(decision))
	}

	cf, err := createFile(folderName + "/" + fileName + ".csv")
	if err != nil {
		return err
	}

	defer cf.Close()

	w := csv.NewWriter(cf)

	err = w.WriteAll(rows)
	if err != nil {
		return err
	}

	return nil
}

// create csv file from operating system
func createFile(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

// generate the necessary headers for csv file
func generateHeaders() []string {
	return []string{
		"Autor",
		"Título",
		"Área do Conhecimento",
		"Tipo Documento",
		"Unidade",
		"Ano",
		"Link",
	}
}

// returns []string that compose the row in the csv file
func generateRow(result Crawler.DocRow) []string {
	return []string{
		result.Name,
		result.Title,
		result.Field,
		result.Type,
		result.Unit,
		result.Year,
		result.Link,
	}
}
