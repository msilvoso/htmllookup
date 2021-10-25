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
func (hp *htmlLookup) headerJson() error {
	// iterate through first csv line -> header (column names)
	var h []columnStruct
	columnLoop:
	for k, column := range hp.header {
		// check if the column should be hidden
		for _, hiddenColumn := range hp.hiddenColumns {
			if k == hiddenColumn {
				continue columnLoop // do not add to the header
			}
		}
		c := columnStruct{Key: column, Sortable: "true"}
		h = append(h, c)
	}
	j, err := json.Marshal(h)
	if err != nil {
		return fmt.Errorf("headerJson : %s", err.Error())
	}
	hp.FieldsJson = string(j)
	return nil
}

// extractColumnNames extracts the column names and populates the header slice
func (hp *htmlLookup) extractColumnNames() error {
	if len(hp.content) == 0 || len(hp.content[0]) == 0 {
		// empty content ?
		return fmt.Errorf("header: content slice seems empty")
	}
	// iterate through first csv line -> header (column names)
	for _, column := range hp.content[0] {
		normColumn, _ := normalizeHeader(column)
		hp.header = append(hp.header, normColumn)
	}
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

// HideColumns defines the columns that should not be shown
// they can still be part of the search
// the arguments are int or string (column index (starting at 0) or name)
func (hp *htmlLookup) HideColumns(columns ...interface{}) error {
	for _, column := range columns {
		switch c := column.(type) {
		case int:
			if c >= len(hp.header) || c < 0 {
				return fmt.Errorf("column index out of bounds\n")
			}
			hp.hiddenColumns = append(hp.hiddenColumns, c)
		case string:
			cIndex, err := hp.columnIndex(c)
			if err != nil {
				return fmt.Errorf("column does not exist\n")
			}
			hp.hiddenColumns = append(hp.hiddenColumns, cIndex)
		default:
			return fmt.Errorf("column is of the wrong type\n")
		}
	}
	return nil
}