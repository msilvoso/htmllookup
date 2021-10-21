package htmllookup

import (
	_ "embed"
	"encoding/csv"
	"os"
	"strings"
	"text/template"
)

//go:embed template.html
var htmlTemplate string

type searchableHtmlPage struct {
	Title            string
	BTableAttributes string
	bTableAttributes []string
	BTableTemplates  string
	bTableTemplates  []string
	ItemLimit        string
	FieldsJson       string
	ItemsJson        string
	SortBy           string
	content          [][]string
	header           []string
	html             string
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
func (hp *searchableHtmlPage) LoadData(fileName string, delimiter rune) error {
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

func (hp searchableHtmlPage) Process() (err error) {
	hp.html, err = hp.generateHtml()
	return err
}

func (hp searchableHtmlPage) generateHtml() (string, error) {
	t := template.Must(template.New("html").Parse(htmlTemplate))
	output := new(strings.Builder)
	err := t.Execute(output, hp)
	if err != nil {
		return "", err
	}
	return output.String(), nil
}
