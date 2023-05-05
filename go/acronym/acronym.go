package acronym

import (
	"regexp"
	"strings"
)

func Abbreviate(s string) (out string) {
	re := regexp.MustCompile(`[( )(\-)(_)]`)
	words := re.Split(s, len(s))

	for _, word := range words {
		if len(word) <= 0 {
			continue
		}
		out += strings.ToUpper(string(word[0]))
	}
	return
}
