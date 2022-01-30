package strings

import (
	"strings"
	"unicode"
)

func countValidWords(sentence string) int {
	words := strings.Fields(sentence)
	cunt := 0
CuntWord:
	for _, word := range words {
		sum := 0
		n := len(word)
		for i := 0; i < n; i++ {
			if unicode.IsLetter(rune(word[i])) {
				continue
			} else if unicode.IsDigit(rune(word[i])) {
				continue CuntWord
			} else if word[i] == '-' {
				if sum == 1 {
					continue CuntWord
				}
				sum++
				if i == 0 || i == n-1 || !unicode.IsLetter(rune(word[i-1])) || !unicode.IsLetter(rune(word[i+1])) {
					continue CuntWord
					break
				}
			} else if i != n-1 {
				continue CuntWord
				break
			}
		}
		cunt++
	}
	return cunt
}
