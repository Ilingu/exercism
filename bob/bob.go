package bob

import (
	"strings"
	"unicode"
)

/* BEST implementation */
// func Hey(remark string) string {
// 	remark = strings.TrimSpace(remark)
// 	if remark == "" {
// 		return "Fine. Be that way!"
// 	}
// 	if strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark {
// 		if strings.HasSuffix(remark, "?") {
// 			return "Calm down, I know what I'm doing!"
// 		}
// 		return "Whoa, chill out!"
// 	}
// 	if strings.HasSuffix(remark, "?") {
// 		return "Sure."
// 	}
// 	return "Whatever."
// }

const (
	sure     = "Sure."
	chill    = "Whoa, chill out!"
	calm     = "Calm down, I know what I'm doing!"
	fine     = "Fine. Be that way!"
	whatever = "Whatever."
)

func Filters(str string) string {
	filters := []string{" ", "\t", "\n", "\r"}
	for _, filtre := range filters {
		str = strings.ReplaceAll(str, filtre, "")
	}
	return str
}

func FilterQuestion(str *string) {
	*str = Filters(*str)
	var newStr string
	for _, char := range *str {
		if unicode.IsLetter(char) || char == '?' {
			newStr += string(char)
		}
	}
	*str = newStr
}

func isUppercase(str string) bool {
	toExclude := 0
	for _, char := range str {
		if !unicode.IsLetter(char) {
			toExclude++
			continue
		}
		if !unicode.IsUpper(char) {
			return false
		}
	}
	return true && len(str)-toExclude > 0
}

func isQuestion(str string) bool {
	return len(str) > 0 && str[len(str)-1] == '?'
}

func Hey(remark string) string {
	if len(Filters(remark)) == 0 {
		return fine
	}

	FilterQuestion(&remark)
	uppercase, question := isUppercase(remark), isQuestion(remark)

	if uppercase && question {
		return calm
	}
	if uppercase {
		return chill
	}
	if question {
		return sure
	}

	return whatever
}
