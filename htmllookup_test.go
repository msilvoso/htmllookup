package htmllookup

import (
	"encoding/csv"
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
	h.DateNow = "2021-11-09 11:53:49"
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
	h.DateNow = "2021-11-09 11:53:49"
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
	h.DateNow = "2021-11-09 11:53:49"
	h.Hover()
	h.Bordered()
	h.Striped()
	err = h.AddOptionToRow("year", OCellIsGreaterOrEqual, 1980, "success")
	if err != nil {
		t.Fatal(err)
	}
	err = h.AddOptionToCell("age", "age", OCellIsLowerOrEqual, 23, "danger")
	if err != nil {
		t.Fatal(err)
	}
	err = h.AddOptionToCell("age", "age", OCellIsGreaterOrEqual, 30, "info")
	if err != nil {
		t.Fatal(err)
	}
	err = h.AddOptionToRow("index", OCellIsLower, 10, "warning")
	if err != nil {
		t.Fatal(err)
	}
	err = h.AddOptionToCell("name", "name", OCellIsGreater, "lll", "primary")
	if err != nil {
		t.Fatal(err)
	}
	err = h.AddOptionToCell("year", "year", OCellIsNotEqual, 2000, "secondary")
	if err != nil {
		t.Fatal(err)
	}
	err = h.HideColumns("index")
	if err != nil {
		t.Fatal(err)
	}
	err = h.SearchableColumns("year", "name", 4)
	if err != nil {
		t.Fatal(err)
	}
	err = h.Process()
	if err != nil {
		t.Fatal(err)
	}
	is := h.Html()
	should, err := ioutil.ReadFile("testdata/refrender2.html")
	if err != nil {
		t.Fatal(err)
	}
	s := string(should)
	if s != is {
		t.Errorf("rendered file different from reference file: \n%s\n", diff(is, s))
	}
}

func TestNewFromData(t *testing.T) {
	fileReader, err := os.Open("testdata/oscar_age_female.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer fileReader.Close()
	c := csv.NewReader(fileReader)
	c.Comma = ','
	c.LazyQuotes = true
	c.TrimLeadingSpace = true
	content, err := c.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	h, err := NewFromData(content)
	h.DateNow = "2021-11-09 11:53:49"
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

func Test_htmlLookup_htmlTemplate(t *testing.T) {
	h := New()
	is := h.htmlTemplate()
	should, err := ioutil.ReadFile("templates/template_en.html")
	if err != nil {
		t.Fatal(err)
	}
	shouldFr, err := ioutil.ReadFile("templates/template_fr.html")
	if err != nil {
		t.Fatal(err)
	}
	s := string(should)
	sFr := string(shouldFr)
	if s != is {
		t.Errorf("rendered file different from reference file: \n%s\n", diff(is, s))
	}
	h.TemplateLanguage = "eng"
	is = h.htmlTemplate()
	if s != is {
		t.Errorf("rendered file different from reference file: \n%s\n", diff(is, s))
	}
	// should default to english
	h.TemplateLanguage = "test"
	is = h.htmlTemplate()
	if s != is {
		t.Errorf("rendered file different from reference file: \n%s\n", diff(is, s))
	}
	// french
	h.TemplateLanguage = "FRançais"
	is = h.htmlTemplate()
	if sFr != is {
		t.Errorf("rendered file different from reference file: \n%s\n", diff(is, sFr))
	}
}

func Test_searchableHtmlPage_AddRelOption(t *testing.T) {
	h, err := NewFromFile("testdata/oscar_age_female.csv", ',')
	if err != nil {
		t.Fatal(err)
	}
	h.DateNow = "2021-11-09 11:53:49"
	h.Hover()
	h.Bordered()
	h.Striped()
	err = h.AddRelOptionToCell("age", "age", OCellIsGreater, "index", true, "danger")
	if err != nil {
		t.Fatal(err)
	}
	err = h.AddRelOptionToCell("name", "name", OCellIsGreaterOrEqual, "name_of_the_movie", true, "warning")
	if err != nil {
		t.Fatal(err)
	}
	err = h.AddRelOptionToRow("index", OCellIsGreater, "age", true, "info")
	if err != nil {
		t.Fatal(err)
	}
	/*err = h.HideColumns("index")
	if err != nil {
		t.Fatal(err)
	}*/
	err = h.SearchableColumns("year", "name", 4)
	if err != nil {
		t.Fatal(err)
	}
	err = h.Process()
	if err != nil {
		t.Fatal(err)
	}
	is := h.Html()
	should, err := ioutil.ReadFile("testdata/refrender3.html")
	if err != nil {
		t.Fatal(err)
	}
	s := string(should)
	if s != is {
		t.Errorf("rendered file different from reference file: \n%s\n", diff(is, s))
	}
}

func Test_searchableHtmlPage_ApplyToColumn(t *testing.T) {
	h, err := NewFromFile("testdata/oscar_age_female.csv", ',')
	if err != nil {
		t.Fatal(err)
	}
	h.DateNow = "2021-11-09 11:53:49"
	h.Hover()
	h.Bordered()
	h.Striped()
	err = h.AddRelOptionToCell("age", "name_of_the_movie", OCellIsGreater, "index", true, "danger")
	if err != nil {
		t.Fatal(err)
	}
	err = h.AddRelOptionToCell("name", "name", OCellIsGreaterOrEqual, "name_of_the_movie", true, "warning")
	if err != nil {
		t.Fatal(err)
	}
	err = h.AddRelOptionToCell("index", "year", OCellIsGreater, "age", true, "info")
	if err != nil {
		t.Fatal(err)
	}
	/*err = h.HideColumns("index")
	if err != nil {
		t.Fatal(err)
	}*/
	err = h.SearchableColumns("year", "name", 4)
	if err != nil {
		t.Fatal(err)
	}
	err = h.Process()
	if err != nil {
		t.Fatal(err)
	}
	is := h.Html()
	should, err := ioutil.ReadFile("testdata/refrender4.html")
	if err != nil {
		t.Fatal(err)
	}
	s := string(should)
	if s != is {
		t.Errorf("rendered file different from reference file: \n%s\n", diff(is, s))
	}
}
