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

type coloringOption struct {
	column    int
	condition int
	compareTo interface{}
	wholeRow  bool
	option    string
}

type searchableHtmlPage struct {
	Title             string
	BTableAttributes  string
	bTableAttributes  []string
	BTableTemplates   string
	bTableTemplates   []string
	ItemLimit         string
	ItemsPerPage      string
	FieldsJson        string
	ItemsJson         string
	SortBy            string
	content           [][]string
	header            []string
	html              string
	coloringOptions   []coloringOption
	hiddenColumns     []int
	searchableColumns []int
}

func New() *searchableHtmlPage {
	// defaults
	s := searchableHtmlPage{Title: "Table Lookup", ItemLimit: "301", ItemsPerPage: "20"}
	return &s
}

func NewFromData(content [][]string) (*searchableHtmlPage, error) {
	s := New()
	s.content = content
	err := s.headerJson()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func NewFromFile(fileName string, delimiter rune) (*searchableHtmlPage, error) {
	s := New()
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
	if err != nil {
		return err
	}
	err = hp.extractColumnNames()
	return err
}

func (hp *searchableHtmlPage) Process() (err error) {
	err = hp.headerJson()
	if err != nil {
		return
	}
	err = hp.itemsJson()
	if err != nil {
		return
	}
	hp.BTableAttributes = strings.Join(hp.bTableAttributes, " ")
	hp.BTableTemplates = strings.Join(hp.bTableTemplates, "")
	hp.html, err = hp.generateHtml()
	return
}

func (hp *searchableHtmlPage) Html() string {
	return hp.html
}

func (hp *searchableHtmlPage) Save(fileName string) error {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(hp.Html())
	return err
}

func (hp *searchableHtmlPage) generateHtml() (string, error) {
	t := template.Must(template.New("html").Parse(htmlTemplate))
	output := new(strings.Builder)
	err := t.Execute(output, hp)
	if err != nil {
		return "", err
	}
	return output.String(), nil
}
