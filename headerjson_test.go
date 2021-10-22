package htmllookup

import "testing"

func Test_searchableHtmlPage_headerJson(t *testing.T) {
	h := New()
	err := h.LoadData("testdata/oscar_age_female.csv", ',')
	if err != nil {
		t.Fatal(err)
	}
	is := h.FieldsJson
	should := `[{"key":"index","sortable":"true"},{"key":"year","sortable":"true"},{"key":"age","sortable":"true"},{"key":"name","sortable":"true"},{"key":"name_of_the_movie","sortable":"true"}]`
	if is != should {
		t.Errorf("generated JSON is not correct: \nis:\n%s\nshould be:\n%s", is, should)
	}
	headerShould := []string{"index", "year", "age", "name", "name_of_the_movie"}
	for k, c := range h.header {
		if c != headerShould[k] {
			t.Errorf("header is not correct: \nis:\n%s\nshould be:\n%s", c, headerShould[k])
		}
	}
}
