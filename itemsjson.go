package htmllookup

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
)

// itemsJson creates the JSON that will be passed to the b-table for the rows
func (hp *searchableHtmlPage) itemsJson() error {
	if len(hp.header) == 0 {
		return fmt.Errorf("the header has to be set (headerJson) before any data can be parsed")
	}
	var jsonMap []map[string]interface{}
	for k, l := range hp.content {
		if k == 0 {
			continue //skip the first line
		}
		line := map[string]interface{}{}
		var searchColumn string
		for hk, cName := range hp.header {
			// value
			line[cName] = l[hk]
			// colorOptions
			rowVariant, cellVariants := hp.colorRowOrCell(hk, l, cName)
			if rowVariant != "" {
				line["_rowVariant"] = rowVariant
			}
			if len(cellVariants) > 0 {
				line["_cellVariants"] = cellVariants
			}
			// add to search
			switch len(hp.searchableColumns) {
			case 0: // if no searchableColumns have been specified all columns should be searchable
				searchColumn += l[hk]
			default:
				var searchable bool
				for _, searchableColumn := range hp.searchableColumns {
					if searchableColumn == hk {
						searchable = true
						break
					}
				}
				if searchable {
					searchColumn += l[hk]
				}
			}
		}
		line["normalized_search_column"], _ = normalize(searchColumn)
		jsonMap = append(jsonMap, line)
	}
	//j,  := json.Marshal(jsonMap)
	jBuf := bytes.NewBuffer([]byte{})
	enc := json.NewEncoder(jBuf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(jsonMap)
	if err != nil {
		return fmt.Errorf("itemsJson: %s", err.Error())
	}
	j := strings.Trim(jBuf.String(), "\n\r ")
	hp.ItemsJson = j
	return nil
}

// colorRowOrCell creates the map that is necessary to set row or cell styles
func (hp *searchableHtmlPage) colorRowOrCell(hk int, l []string, cName string) (rowVariant string, cellVariants map[string]string) {
	// there can only be one rowVariant but multiple cellVariants
	for _, option := range hp.coloringOptions {
		if option.column == hk {
			value := l[hk]
			if checks, _ := checkCondition(value, option.condition, option.compareTo); checks {
				if option.wholeRow {
					rowVariant = option.option
					continue
				}
				if len(cellVariants) == 0 {
					cellVariants = map[string]string{}
				}
				cellVariants[cName] = option.option
			}
		}
	}
	return rowVariant, cellVariants
}

// SearchableColumns defines the columns that should be part of the search index
// the arguments are int or string (column index (starting at 0) or name)
func (hp *searchableHtmlPage) SearchableColumns(columns ...interface{}) error {
	for _, column := range columns {
		switch c := column.(type) {
		case int:
			if c >= len(hp.header) || c < 0 {
				return fmt.Errorf("column index out of bounds\n")
			}
			hp.searchableColumns = append(hp.searchableColumns, c)
		case string:
			cIndex, err := hp.columnIndex(c)
			if err != nil {
				return fmt.Errorf("column does not exist\n")
			}
			hp.searchableColumns = append(hp.searchableColumns, cIndex)
		default:
			return fmt.Errorf("column is of the wrong type\n")
		}
	}
	return nil
}

// normalize removes spaces, transliterates special characters, converts to lowercase
func normalize(input string) (string, error) {
	t := transform.Chain(cases.Lower(language.Und), runes.Remove(runes.In(unicode.White_Space)), norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	normalized, _, err := transform.String(t, input)
	if err != nil {
		return input, err
	}
	return normalized, nil
}
