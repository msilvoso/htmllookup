package htmllookup

import (
	"embed"
	_ "embed"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"
)

//go:embed templates
var templateStore embed.FS

type coloringOption struct {
	column             int         // the column (numeric value) to be compared
	applyOptionTo      int         // the column the style has to be applied to
	relativeComparison bool        // compare to a column instead of a direct value
	numericComparison  bool        // compare two columns numerically (only used when relativeComparison is true)
	condition          int         // condition code -> see OCell... constants
	compareTo          interface{} // value to be compared to
	compareToColumn    int         // column to be compared to
	wholeRow           bool        // color the whole row (not just the cell)
	option             string      // coloring class to add -> see bootstrapvue classes
}

type htmlLookup struct {
	Title                 string
	TemplateLanguage      string
	BTableAttributes      string
	bTableAttributes      []string
	BTableTemplates       string
	bTableTemplates       []string
	ItemLimit             string
	ItemsPerPage          string
	FieldsJson            string
	ItemsJson             string
	SortBy                string
	content               [][]string
	header                []string
	html                  string
	coloringOptions       []coloringOption
	hiddenColumns         []int
	searchableColumns     []int
	notHtmlEscapedColumns []int
	DateNow               string
}

// New is the simple factory for the htmlLookup struct
// you will still need to load some data
func New() *htmlLookup {
	// defaults
	s := htmlLookup{Title: "Table Lookup", ItemLimit: "301", ItemsPerPage: "20", TemplateLanguage: "en", DateNow: time.Now().Format(`2006-01-02 15:04:05`)}
	return &s
}

// NewFromData instantiates a new htmlLookup struct preloading it with the content of a slice
func NewFromData(content [][]string) (*htmlLookup, error) {
	s := New()
	s.content = content
	err := s.extractColumnNames()
	if err != nil {
		return nil, err
	}
	return s, nil
}

// NewFromFile instantiates a new htmlLookup struct preloading it with the content of a csv file
func NewFromFile(fileName string, delimiter rune) (*htmlLookup, error) {
	s := New()
	err := s.LoadData(fileName, delimiter)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// LoadData loads data from a csv file
func (hp *htmlLookup) LoadData(fileName string, delimiter rune) error {
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

func (hp *htmlLookup) Process() (err error) {
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

func (hp *htmlLookup) Html() string {
	return hp.html
}

func (hp *htmlLookup) Save(fileName string) error {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(hp.Html())
	return err
}

func (hp *htmlLookup) generateHtml() (string, error) {
	t := template.Must(template.New("html").Parse(hp.htmlTemplate()))
	output := new(strings.Builder)
	err := t.Execute(output, hp)
	if err != nil {
		return "", err
	}
	return output.String(), nil
}

func (hp *htmlLookup) htmlTemplate() string {
	templt, err := templateStore.ReadFile(fmt.Sprintf("templates/template_%s.html", strings.ToLower(hp.TemplateLanguage[0:2])))
	if err != nil {
		templt, _ = templateStore.ReadFile("templates/template_en.html")
	}
	return string(templt)
}
