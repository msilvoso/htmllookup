package htmllookup

import "testing"

func Test_checkCondition(t *testing.T) {
	count := 1
	if ok, _ := checkCondition("23", OCellIsGreater, int64(24)); ok {
		t.Errorf("error on check %d\n", count)
	}
	count++
	if ok, _ := checkCondition("23", OCellIsLowerOrEqual, int64(24)); !ok {
		t.Errorf("error on check %d\n", count)
	}
	count++
	if ok, _ := checkCondition("", OCellHasValue|OCellIsEqual, ""); ok {
		t.Errorf("error on check %d\n", count)
	}
	count++
	if ok, _ := checkCondition("AB", OCellHasValue|OCellIsGreater, "AA"); !ok {
		t.Errorf("error on check %d\n", count)
	}
	count++
	if ok, _ := checkCondition("", OCellIsEqual, ""); !ok {
		t.Errorf("error on check %d\n", count)
	}
	count++
	// this should fail -> no a number
	if ok, _ := checkCondition("A", OCellIsEqual, 33.33); ok {
		t.Errorf("error on check %d\n", count)
	}
	count++
	if ok, _ := checkCondition("33.33", OCellIsEqual, 33.33); !ok {
		t.Errorf("error on check %d\n", count)
	}
	count++
	if ok, _ := checkCondition("33.33", OCellIsGreaterOrEqual, 33.33); !ok {
		t.Errorf("error on check %d\n", count)
	}
	count++
}
