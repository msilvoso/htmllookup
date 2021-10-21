package htmllookup

import (
	"encoding/json"
	"fmt"
)

type columnStruct struct {
	Key      string `json:"key"`
	Sortable string `json:"sortable"`
}

// headerJson creates the JSON that defines the table header
func (hp searchableHtmlPage) headerJson() (string, error) {
	if len(hp.content) == 0 || len(hp.content[0]) == 0 {
		// empty content ?
		return "", fmt.Errorf("headerJson: content slice seems empty")
	}
	// iterate through first csv line -> header (column names)
	var h []columnStruct
	for _, column := range hp.content[0] {
		c := columnStruct{Key: column, Sortable: "true"}
		h = append(h, c)
	}
	j, err := json.Marshal(h)
	if err != nil {
		return "", fmt.Errorf("headerJson : %s", err.Error())
	}
	return string(j), nil
}
