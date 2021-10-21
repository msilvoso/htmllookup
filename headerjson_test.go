package htmllookup

import "testing"

func Test_searchableHtmlPage_headerJson(t *testing.T) {
	h := New()
	err := h.LoadData("testdata/oscar_age_female.csv", ',')
	if err != nil {
		t.Fatal(err)
	}
	err = h.headerJson()
	if err != nil {
		t.Fatal(err)
	}
	is := h.FieldsJson
	should := `[{"key":"Index","sortable":"true"},{"key":"Year","sortable":"true"},{"key":"Age","sortable":"true"},{"key":"Name","sortable":"true"},{"key":"Movie","sortable":"true"}]`
	if is != should {
		t.Errorf("generated JSON is not correct: \nis:\n%s\nshould be:\n%s", is, should)
	}
	headerShould := []string{"Index", "Year", "Age", "Name", "Movie"}
	for k, c := range h.header {
		if c != headerShould[k] {
			t.Errorf("header is not correct: \nis:\n%s\nshould be:\n%s", c, headerShould[k])
		}
	}
}
