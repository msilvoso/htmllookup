package htmllookup

import (
	"encoding/json"
	"fmt"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"unicode"
)

// itemsJson creates the JSON that will be passed to the b-table for the rows
func (hp *searchableHtmlPage) itemsJson() error {
	if len(hp.header) == 0 {
		return fmt.Errorf("the header has to be set (headerJson) before any data can be parsed")
	}
	var jsonMap []map[string]string
	for k, l := range hp.content {
		if k == 0 {
			continue //skip the first line
		}
		line := map[string]string{}
		var searchColumn string
		for hk, cName := range hp.header {
			line[cName] = l[hk]
			searchColumn += l[hk]
		}
		line["normalized_search_column"], _ = normalize(searchColumn)
		jsonMap = append(jsonMap, line)
	}
	j, err := json.Marshal(jsonMap)
	if err != nil {
		return fmt.Errorf("itemsJson: %s", err.Error())
	}
	hp.ItemsJson = string(j)
	return nil
}

//
func normalize(input string) (string, error) {
	t := transform.Chain(runes.Remove(runes.In(unicode.White_Space)), norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	normalized, _, err := transform.String(t, input)
	if err != nil {
		return input, err
	}
	return normalized, nil
}