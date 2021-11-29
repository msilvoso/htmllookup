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
}

func Test_checkRelCondition(t *testing.T) {
	count := 1
	if ok, _ := checkRelCondition("23", "24", OCellIsGreater, true); ok {
		t.Errorf("error on check %d\n", count)
	}
	count++
	if ok, _ := checkRelCondition("23", "24", OCellIsLowerOrEqual, true); !ok {
		t.Errorf("error on check %d\n", count)
	}
	count++
	if ok, _ := checkRelCondition("", "", OCellHasValue|OCellIsEqual, false); ok {
		t.Errorf("error on check %d\n", count)
	}
	count++
	if ok, _ := checkRelCondition("AB", "AA", OCellHasValue|OCellIsGreater, false); !ok {
		t.Errorf("error on check %d\n", count)
	}
	count++
	if ok, _ := checkRelCondition("", "", OCellIsEqual, false); !ok {
		t.Errorf("error on check %d\n", count)
	}
	count++
	// this should fail -> no a number
	if ok, _ := checkRelCondition("A", "33.33", OCellIsEqual, true); ok {
		t.Errorf("error on check %d\n", count)
	}
	count++
	if ok, _ := checkRelCondition("33.33", "33.33", OCellIsEqual, true); !ok {
		t.Errorf("error on check %d\n", count)
	}
	count++
	if ok, _ := checkRelCondition("33.33", "33.33", OCellIsGreaterOrEqual, true); !ok {
		t.Errorf("error on check %d\n", count)
	}
	count++
	if ok, _ := checkRelCondition("A", "A", OCellIsGreaterOrEqual, true); !ok {
		t.Errorf("error on check %d\n", count)
	}
}
