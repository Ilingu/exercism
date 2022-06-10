package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}

	var sum int
	for i, char := range id {
		if !unicode.IsDigit(char) {
			return false
		}

		num := int(char - '0')
		if (len([]rune(id))%2 == 0 && (i+1)%2 == 1) || len([]rune(id))%2 == 1 && (i+1)%2 == 0 {
			if num*2 > 9 {
				sum += ((num * 2) - 9)
				continue
			}
			sum += (num * 2)
			continue
		}
		sum += num
	}

	return sum%10 == 0
}
