package htmllookup

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func Test_searchableHtmlPage_LoadData(t *testing.T) {
	h := New()
	err := h.LoadData("testdata/oscar_age_female.csv", ',')
	if err != nil {
		t.Fatal(err)
	}
	testMap := map[string]string{
		h.content[82][3]: "Kate Winslet",
		h.content[70][4]: "Fargo",
		h.content[42][0]: "42",
		h.content[42][1]: "1969",
	}
	for is, wanted := range testMap {
		if is != wanted {
			t.Errorf("%s is not %s", is, wanted)
		}
	}
}

func Test_searchableHtmlPage_GenerateHtml1(t *testing.T) {
	h := New()
	h.ItemLimit = "#######ITEMLIMIT***REPLACEMENT#######"
	h.BTableAttributes = "#######BTABLEATTRIBUTES***REPLACEMENT#######"
	h.BTableTemplates = "#######BTABLESTEMPLATES***REPLACEMENT#######"
	h.FieldsJson = "#######FIELDSJSON***REPLACEMENT#######"
	h.ItemsJson = "#######ITEMSJSON***REPLACEMENT#######"
	h.SortBy = "#######SORTBY***REPLACEMENT#######"
	h.Title = "#######TITLE***REPLACEMENT#######"
	h.ItemsPerPage = "#######PERPAGE***REPLACEMENT#######"
	is, err := h.generateHtml()
	if err != nil {
		t.Fatal(err)
	}
	f, err := os.Open("testdata/process1test.html")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	should, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	difference := diff(string(should), is)
	if difference != "" {
		t.Errorf("There is a problem : \n%s\n", difference)
	}
}

func Test_searchableHtmlPage_Html(t *testing.T) {
	h, err := NewFromFile("testdata/oscar_age_female.csv", ',')
	if err != nil {
		t.Fatal(err)
	}
	err = h.Process()
	if err != nil {
		t.Fatal(err)
	}
	is := h.Html()
	should, err := ioutil.ReadFile("testdata/refrender.html")
	if err != nil {
		t.Fatal(err)
	}
	s := string(should)
	if s != is {
		t.Errorf("rendered file different from reference file: \n%s\n", diff(is, s))
	}
}

func Test_searchableHtmlPage_Save(t *testing.T) {
	h, err := NewFromFile("testdata/oscar_age_female.csv", ',')
	if err != nil {
		t.Fatal(err)
	}
	h.Hover()
	h.Bordered()
	h.Striped()
	for k := range h.content {
		if k == 0 {
			continue
		}
		h.content[k][1] = fmt.Sprintf("<b>%s</b>", h.content[k][1])
	}
	err = h.RenderInRawHtml("year")
	if err != nil {
		t.Fatal(err)
	}
	err = h.Process()
	if err != nil {
		t.Fatal(err)
	}
	err = h.Save("/mnt/test.html")
	if err != nil {
		t.Fatal(err)
	}
}