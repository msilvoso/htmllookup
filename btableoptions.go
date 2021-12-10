package htmllookup

import "fmt"

// Hover adds the "hover" attribute to the btable
func (hp *htmlLookup) Hover() {
	hp.AddBTableAttribute("hover")
}

// Bordered adds the "bordered" attribute to the btable
func (hp *htmlLookup) Bordered() {
	hp.AddBTableAttribute("bordered")
}

// Striped adds the "striped" attribute to the btable
func (hp *htmlLookup) Striped() {
	hp.AddBTableAttribute("striped")
}

// AddBTableAttribute adds an attribute to the bTable
func (hp *htmlLookup) AddBTableAttribute(attribute string) {
	hp.bTableAttributes = append(hp.bTableAttributes, attribute)
}

// RenderInRawHtml adds template to bTable to render the column values in raw html
// column can either be an int, string or a []byte
func (hp *htmlLookup) RenderInRawHtml(column interface{}) (err error) {
	var cName string
	var cIndex int
	switch t := column.(type) {
	case int:
		cName, err = hp.columnName(t)
		if err != nil {
			return err
		}
		cIndex = t
	case string:
		cName = t
		cIndex, err = hp.columnIndex(cName)
		if err != nil {
			return err
		}
	case []byte:
		cName = string(t)
		cIndex, err = hp.columnIndex(cName)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("column argument has to be either int, string or []byte")
	}
	hp.AddBTableTemplate(fmt.Sprintf("<template #cell(%s)=\"data\"><span v-html=\"data.value\"></span></template>", cName))
	hp.notHtmlEscapedColumns = append(hp.notHtmlEscapedColumns, cIndex)
	return
}

// AddBTableTemplate adds template to the bTable
// I do not know any other application than to render in raw html
func (hp *htmlLookup) AddBTableTemplate(template string) {
	hp.bTableTemplates = append(hp.bTableTemplates, template)
}

// columnName returns the column name
func (hp *htmlLookup) columnName(col int) (string, error) {
	if col > len(hp.header)-1 {
		return "", fmt.Errorf("the column with index %d does not exist", col)
	}
	return hp.header[col], nil
}

// columnName returns the column name
func (hp *htmlLookup) columnIndex(col string) (int, error) {
	for k, c := range hp.header {
		if col == c {
			return k, nil
		}
	}
	return -1, fmt.Errorf("column not found")
}
