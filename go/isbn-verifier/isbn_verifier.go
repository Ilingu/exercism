package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	isbnChar := strings.Split(isbn, "")

	if len(isbnChar) != 10 {
		return false
	}

	var sum int
	for i, char := range isbnChar {
		digit, err := strconv.Atoi(char)

		if char == "X" {
			digit = 10
		}
		if err != nil && (char != "X" || i != len(isbnChar)-1) {
			return false
		}

		sum += digit * (10 - i)
	}

	return sum%11 == 0
}
