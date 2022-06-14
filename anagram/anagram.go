package anagram

import (
	"strings"
	"unicode"
)

func IndexStr(str string) map[rune]int {
	out := make(map[rune]int)
	for _, char := range str {
		out[unicode.ToLower(char)]++
	}
	return out
}

func CheckCandidate(candidate string, ch chan bool) {
	if strings.EqualFold(candidate, Gsubject) {
		ch <- false
	}

	indexedCandidate := IndexStr(candidate)

	for letter, frequency := range indexedCandidate {
		if freq, found := indexedSubject[letter]; !found || freq != frequency {
			ch <- false
		}
	}
	ch <- true
}

var indexedSubject map[rune]int
var Gsubject string

func Detect(subject string, candidates []string) []string {
	indexedSubject = IndexStr(subject)
	Gsubject = subject

	validCandidates := make([]string, 0)
	for _, candidate := range candidates {
		res := make(chan bool)
		go CheckCandidate(candidate, res) // Using Goroutines (Faster but use more performance)
		valid := <-res

		if valid {
			validCandidates = append(validCandidates, candidate)
		}
	}
	return validCandidates
}
