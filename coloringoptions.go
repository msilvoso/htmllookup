package htmllookup

import (
	"fmt"
	"strconv"
	"strings"
)

// powers of two
const (
	OCellHasValue         = 1 << iota
	OCellIsEqual          = 1 << iota
	OCellIsGreater        = 1 << iota
	OCellIsLower          = 1 << iota
	OCellIsGreaterOrEqual = 1 << iota
	OCellIsLowerOrEqual   = 1 << iota
)

// checkCondition checks the content of the cell against a value
// the comparison can be numeric or string
// use int64 and float64 only
func checkCondition(value string, condition int, compareTo interface{}) (bool, error) {
	// check if the cell is empty
	if condition&OCellHasValue != 0 && value == "" {
		return false, nil
	}
	switch v := compareTo.(type) {
	case int64:
		c, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			break
		}
		switch {
		case condition&OCellIsEqual != 0:
			if c == v {
				return true, nil
			}
		case condition&OCellIsGreater != 0:
			if c > v {
				return true, nil
			}
		case condition&OCellIsLower != 0:
			if c < v {
				return true, nil
			}
		case condition&OCellIsGreaterOrEqual != 0:
			if c >= v {
				return true, nil
			}
		case condition&OCellIsLowerOrEqual != 0:
			if c <= v {
				return true, nil
			}
		}
	case float64:
		c, err := strconv.ParseFloat(value, 64)
		if err != nil {
			break
		}
		switch {
		case condition&OCellIsEqual != 0:
			if c == v {
				return true, nil
			}
		case condition&OCellIsGreater != 0:
			if c > v {
				return true, nil
			}
		case condition&OCellIsLower != 0:
			if c < v {
				return true, nil
			}
		case condition&OCellIsGreaterOrEqual != 0:
			if c >= v {
				return true, nil
			}
		case condition&OCellIsLowerOrEqual != 0:
			if c <= v {
				return true, nil
			}
		}
	case string:
		switch {
		case condition&OCellIsEqual != 0:
			if strings.ToLower(value) == strings.ToLower(v) {
				return true, nil
			}
		case condition&OCellIsGreater != 0:
			if strings.ToLower(value) > strings.ToLower(v) {
				return true, nil
			}
		case condition&OCellIsLower != 0:
			if strings.ToLower(value) < strings.ToLower(v) {
				return true, nil
			}
		case condition&OCellIsGreaterOrEqual != 0:
			if strings.ToLower(value) >= strings.ToLower(v) {
				return true, nil
			}
		case condition&OCellIsLowerOrEqual != 0:
			if strings.ToLower(value) <= strings.ToLower(v) {
				return true, nil
			}
		}
	default:
		return false, fmt.Errorf("comparison value should be int64, float64 or string\n")
	}
	return false, nil
}

func (hp *searchableHtmlPage) AddOption(column interface{}, condition int, compareTo interface{}, wholeRow bool, option string) (err error) {
	cOption := coloringOption{condition: condition, wholeRow: wholeRow, option: option}
	switch col := column.(type) {
	case int:
		if col >= len(hp.header) {
			return fmt.Errorf("column index out of bounds\n")
		}
		cOption.column = col
	case string:
		cOption.column, err = hp.columnIndex(col)
		if err != nil {
			return fmt.Errorf("column does not exist\n")
		}
	default:
		return fmt.Errorf("column is of the wrong type\n")
	}
	switch cp := compareTo.(type) {
	case int:
		cOption.compareTo = int64(cp)
	case int8:
		cOption.compareTo = int64(cp)
	case int16:
		cOption.compareTo = int64(cp)
	case int32:
		cOption.compareTo = int64(cp)
	case float32:
		cOption.compareTo = float64(cp)
	default:
		cOption.compareTo = cp
	}
	hp.coloringOptions = append(hp.coloringOptions, cOption)
	return
}
