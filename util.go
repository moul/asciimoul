package asciimoul

import "strings"

func reverseMap(r rune) rune {
	mapping := map[rune]rune{
		'<':  '>',
		'>':  '<',
		')':  '(',
		'(':  ')',
		'/':  '\\',
		'\\': '/',
	}
	if match, found := mapping[r]; found {
		return match
	}
	return r
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return strings.Map(reverseMap, string(runes))
}
