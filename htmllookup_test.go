package htmllookup

import (
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
