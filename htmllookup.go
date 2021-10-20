package htmllookup

import (
	_ "embed"
	"encoding/csv"
	"os"
)

//go:embed template.html
var htmlTemplate string

type searchableHtmlPage struct {
	Title string
	BTableAttributes string
	bTableAttributes []string
	BTableTemplates string
	bTableTemplates []string
	ItemLimit string
	FieldsJson string
	ItemsJson string
	content [][]string
}

func New() *searchableHtmlPage {
	return new(searchableHtmlPage)
}

func NewFromData(content [][]string) *searchableHtmlPage {
	return &searchableHtmlPage{content: content}
}

func NewFromFile(fileName string, delimiter rune) (*searchableHtmlPage, error) {
	s := new(searchableHtmlPage)
	err := s.LoadData(fileName, delimiter)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// LoadData loads data from a csv file
func (hp *searchableHtmlPage) LoadData(fileName string, delimiter rune) error{
	fileReader, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer fileReader.Close()
	c := csv.NewReader(fileReader)
	c.Comma = delimiter
	c.LazyQuotes = true
	c.TrimLeadingSpace = true
	hp.content, err = c.ReadAll()
	return err
}


