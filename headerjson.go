package htmllookup

import (
	"encoding/json"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"unicode"
)

type columnStruct struct {
	Key      string `json:"key"`
	Sortable string `json:"sortable"`
}

// headerJson creates the JSON that defines the table header
func (hp *searchableHtmlPage) headerJson() error {
	if len(hp.content) == 0 || len(hp.content[0]) == 0 {
		// empty content ?
		return fmt.Errorf("headerJson: content slice seems empty")
	}
	// iterate through first csv line -> header (column names)
	var h []columnStruct
	for _, column := range hp.content[0] {
		normColumn, _ := normalizeHeader(column)
		hp.header = append(hp.header, normColumn)
		c := columnStruct{Key: normColumn, Sortable: "true"}
		h = append(h, c)
	}
	j, err := json.Marshal(h)
	if err != nil {
		return fmt.Errorf("headerJson : %s", err.Error())
	}
	hp.FieldsJson = string(j)
	return nil
}

// normalizeHeader converts spaces to _, transliterates special characters, converts to lowercase
func normalizeHeader(input string) (string, error) {
	replaceWhiteSpacesByUnderScores := runes.Map(func(r rune) rune {
		if unicode.Is(unicode.White_Space, r) {
			return '_'
		}
		return r
	})
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC, cases.Lower(language.Und), replaceWhiteSpacesByUnderScores)
	normalized, _, err := transform.String(t, input)
	if err != nil {
		return input, err
	}
	return normalized, nil
}
