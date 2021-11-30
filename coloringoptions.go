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
	OCellIsNotEqual       = 1 << iota
)

func (hp *htmlLookup) AddOptionToRow(column interface{}, condition int, compareTo interface{}, option string) error {
	return hp.addOption(column, column, condition, compareTo, true, option)
}

func (hp *htmlLookup) AddOptionToCell(column interface{}, applyOptionTo interface{}, condition int, compareTo interface{}, option string) error {
	return hp.addOption(column, applyOptionTo, condition, compareTo, false, option)
}

func (hp *htmlLookup) AddRelOptionToRow(column interface{}, condition int, compareToColumn interface{}, numericComparison bool, option string) error {
	return hp.addRelOption(column, column, condition, compareToColumn, numericComparison, true, option)
}

func (hp *htmlLookup) AddRelOptionToCell(column interface{}, applyOptionTo interface{}, condition int, compareToColumn interface{}, numericComparison bool, option string) error {
	return hp.addRelOption(column, applyOptionTo, condition, compareToColumn, numericComparison, false, option)
}

// checkCondition checks the content of the cell against a value
// the comparison can be numeric or string
// use int64 and float64 only
func checkCondition(value string, condition int, compareTo interface{}) (bool, error) {
	// check if the cell is empty
	if condition&OCellHasValue != 0 && value == "" {
		return false, nil
	}
	switch cptTo := compareTo.(type) {
	case int64:
		val, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			break
		}
		switch {
		case condition&OCellIsEqual != 0:
			if val == cptTo {
				return true, nil
			}
		case condition&OCellIsGreater != 0:
			if val > cptTo {
				return true, nil
			}
		case condition&OCellIsLower != 0:
			if val < cptTo {
				return true, nil
			}
		case condition&OCellIsGreaterOrEqual != 0:
			if val >= cptTo {
				return true, nil
			}
		case condition&OCellIsLowerOrEqual != 0:
			if val <= cptTo {
				return true, nil
			}
		case condition&OCellIsNotEqual != 0:
			if val != cptTo {
				return true, nil
			}
		}
	case float64:
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			break
		}
		switch {
		case condition&OCellIsEqual != 0:
			if val == cptTo {
				return true, nil
			}
		case condition&OCellIsGreater != 0:
			if val > cptTo {
				return true, nil
			}
		case condition&OCellIsLower != 0:
			if val < cptTo {
				return true, nil
			}
		case condition&OCellIsGreaterOrEqual != 0:
			if val >= cptTo {
				return true, nil
			}
		case condition&OCellIsLowerOrEqual != 0:
			if val <= cptTo {
				return true, nil
			}
		case condition&OCellIsNotEqual != 0:
			if val != cptTo {
				return true, nil
			}
		}
	case string:
		switch {
		case condition&OCellIsEqual != 0:
			if strings.ToLower(value) == strings.ToLower(cptTo) {
				return true, nil
			}
		case condition&OCellIsGreater != 0:
			if strings.ToLower(value) > strings.ToLower(cptTo) {
				return true, nil
			}
		case condition&OCellIsLower != 0:
			if strings.ToLower(value) < strings.ToLower(cptTo) {
				return true, nil
			}
		case condition&OCellIsGreaterOrEqual != 0:
			if strings.ToLower(value) >= strings.ToLower(cptTo) {
				return true, nil
			}
		case condition&OCellIsLowerOrEqual != 0:
			if strings.ToLower(value) <= strings.ToLower(cptTo) {
				return true, nil
			}
		case condition&OCellIsNotEqual != 0:
			if strings.ToLower(value) != strings.ToLower(cptTo) {
				return true, nil
			}
		}
	default:
		return false, fmt.Errorf("comparison value should be int64, float64 or string\n")
	}
	return false, nil
}

func (hp *htmlLookup) addOption(column interface{}, applyOptionTo interface{}, condition int, compareTo interface{}, wholeRow bool, option string) (err error) {
	cOption := coloringOption{condition: condition, wholeRow: wholeRow, option: option}
	switch col := column.(type) {
	case int:
		if col >= len(hp.header) || col < 0 {
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
	switch col := applyOptionTo.(type) {
	case int:
		if col >= len(hp.header) || col < 0 {
			return fmt.Errorf("applyOptionTo index out of bounds\n")
		}
		cOption.applyOptionTo = col
	case string:
		cOption.applyOptionTo, err = hp.columnIndex(col)
		if err != nil {
			return fmt.Errorf("applyOptionTo does not exist\n")
		}
	default:
		return fmt.Errorf("applyOptionTo is of the wrong type\n")
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

// checkRelCondition checks the content of the cell against a value in another column
// the comparison can be numeric or string according to numericComparison
func checkRelCondition(value string, compareTo string, condition int, numericComparison bool) (bool, error) {
	// check if the cell is empty
	if condition&OCellHasValue != 0 && value == "" {
		return false, nil
	}
	switch numericComparison {
	case true:
		var notNumeric bool
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			notNumeric = true
		}
		cptTo, err := strconv.ParseFloat(compareTo, 64)
		if err != nil {
			notNumeric = true
		}
		switch {
		case notNumeric:
			break
		case condition&OCellIsEqual != 0:
			if val == cptTo {
				return true, nil
			}
			return false, nil // because of the fallthrough you have to exit
		case condition&OCellIsGreater != 0:
			if val > cptTo {
				return true, nil
			}
			return false, nil // because of the fallthrough you have to exit
		case condition&OCellIsLower != 0:
			if val < cptTo {
				return true, nil
			}
			return false, nil // because of the fallthrough you have to exit
		case condition&OCellIsGreaterOrEqual != 0:
			if val >= cptTo {
				return true, nil
			}
			return false, nil // because of the fallthrough you have to exit
		case condition&OCellIsLowerOrEqual != 0:
			if val <= cptTo {
				return true, nil
			}
			return false, nil // because of the fallthrough you have to exit
		case condition&OCellIsNotEqual != 0:
			if val != cptTo {
				return true, nil
			}
			return false, nil // because of the fallthrough you have to exit
		}
		fallthrough
	default:
		switch {
		case condition&OCellIsEqual != 0:
			if strings.ToLower(value) == strings.ToLower(compareTo) {
				return true, nil
			}
		case condition&OCellIsGreater != 0:
			if strings.ToLower(value) > strings.ToLower(compareTo) {
				return true, nil
			}
		case condition&OCellIsLower != 0:
			if strings.ToLower(value) < strings.ToLower(compareTo) {
				return true, nil
			}
		case condition&OCellIsGreaterOrEqual != 0:
			if strings.ToLower(value) >= strings.ToLower(compareTo) {
				return true, nil
			}
		case condition&OCellIsLowerOrEqual != 0:
			if strings.ToLower(value) <= strings.ToLower(compareTo) {
				return true, nil
			}
		case condition&OCellIsNotEqual != 0:
			if strings.ToLower(value) != strings.ToLower(compareTo) {
				return true, nil
			}
		}
	}
	return false, nil
}

func (hp *htmlLookup) addRelOption(column interface{}, applyOptionTo interface{}, condition int, compareToColumn interface{}, numericComparison bool, wholeRow bool, option string) (err error) {
	cOption := coloringOption{
		condition:          condition,
		wholeRow:           wholeRow,
		option:             option,
		numericComparison:  numericComparison,
		relativeComparison: true,
	}
	switch col := column.(type) {
	case int:
		if col >= len(hp.header) || col < 0 {
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
	switch col := applyOptionTo.(type) {
	case int:
		if col >= len(hp.header) || col < 0 {
			return fmt.Errorf("applyOptionTo index out of bounds\n")
		}
		cOption.applyOptionTo = col
	case string:
		cOption.applyOptionTo, err = hp.columnIndex(col)
		if err != nil {
			return fmt.Errorf("applyOptionTo does not exist\n")
		}
	default:
		return fmt.Errorf("applyOptionTo is of the wrong type\n")
	}
	switch col := compareToColumn.(type) {
	case int:
		if col >= len(hp.header) || col < 0 {
			return fmt.Errorf("column index out of bounds\n")
		}
		cOption.compareToColumn = col
	case string:
		cOption.compareToColumn, err = hp.columnIndex(col)
		if err != nil {
			return fmt.Errorf("compareToColumn does not exist\n")
		}
	default:
		return fmt.Errorf("compareToColumn is of the wrong type\n")
	}
	hp.coloringOptions = append(hp.coloringOptions, cOption)
	return
}
