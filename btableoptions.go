package htmllookup

import "fmt"

// Hover adds the "hover" attribute to the btable
func (hp *searchableHtmlPage) Hover() {
	hp.AddBTableAttribute("hover")
}

// Bordered adds the "bordered" attribute to the btable
func (hp *searchableHtmlPage) Bordered() {
	hp.AddBTableAttribute("bordered")
}

// Striped adds the "striped" attribute to the btable
func (hp *searchableHtmlPage) Striped() {
	hp.AddBTableAttribute("striped")
}

// AddBTableAttribute adds an attribute to the bTable
func (hp *searchableHtmlPage) AddBTableAttribute(attribute string) {
	hp.bTableAttributes = append(hp.bTableAttributes, attribute)
}

// RenderInRawHtml adds template to bTable to render the column values in raw html
// column can either be an int, string or a []byte
func (hp *searchableHtmlPage) RenderInRawHtml(column interface{}) (err error){
	var cName string
	switch t := column.(type) {
	case int:
		cName, err = hp.columnName(t)
		if err != nil {
			return err
		}
	case string:
		cName = t
	case []byte:
		cName = string(t)
	default:
		return fmt.Errorf("column argument has to be either int, string or []byte")
	}
	hp.AddBTableTemplate(fmt.Sprintf("<template #cell(%s)=\"data\"><span v-html=\"data.value\"></span></template>", cName))
	return
}

// AddBTableTemplate adds template to the bTable
// I do not know any other application than to render in raw html
func (hp *searchableHtmlPage) AddBTableTemplate(template string) {
	hp.bTableTemplates = append(hp.bTableTemplates, template)
}

// columnName returns the column name
func (hp *searchableHtmlPage) columnName(col int) (string, error) {
	if col > len(hp.header) - 1 {
		return "", fmt.Errorf("the column with index %d does not exist", col)
	}
	return hp.header[col], nil
}
