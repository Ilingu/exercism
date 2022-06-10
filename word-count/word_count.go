package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int
type Found struct {
	Exists bool
	Index  int
}

func removeDuplicates(arr [][]string) []string {
	occurred := map[string]Found{}
	result := []string{}
	for i, words := range arr {
		for _, word := range words {
			if f, ok := occurred[word]; !ok || f.Index == i {
				occurred[word] = Found{Exists: true, Index: i}
				result = append(result, word)
			}
		}
	}
	return result
}

func WordCount(phrase string) Frequency {
	reg := regexp.MustCompile("[^a-zA-Z0-9']+")

	phrase = reg.ReplaceAllString(phrase, " ")
	phrase = strings.Trim(phrase, " ")
	separators := []string{" ", "\n", "\t", ","}
	WordsRes := [][]string{}
	maxResLen := 0
	for _, sep := range separators {
		res := strings.Split(phrase, sep)
		if len(res) <= 1 && len(res) < maxResLen {
			continue
		}
		if len(res) > maxResLen {
			maxResLen = len(res)
		}
		WordsRes = append(WordsRes, res)
	}
	Words := removeDuplicates(WordsRes)

	out := make(Frequency)
	reg = regexp.MustCompile(`'\w+'`)
	for _, word := range Words {
		iWord := strings.ToLower(word)
		if reg.MatchString(iWord) {
			iWord = strings.ReplaceAll(iWord, "'", "")
		}

		out[iWord]++
	}
	return out
}
