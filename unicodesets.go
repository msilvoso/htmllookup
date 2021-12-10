package htmllookup

import "unicode"

type nonWordCharsSet struct {
}

func (n nonWordCharsSet) Contains(r rune) bool {
	if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsSpace(r) && r != '_' {
		return true
	}
	return false
}

type removeHtmlSpecialChars struct {
}

func (n removeHtmlSpecialChars) Contains(r rune) bool {
	if r == '<' || r == '>' || r == '"' {
		return true
	}
	return false
}
