package phonenumber

import (
	"errors"
	"fmt"
	"unicode"
)

func CleanUpNumber(num string) (string, error) {
	var CleanedUp string
	for i, ch := range num {
		if unicode.IsDigit(ch) && !((i == 0 || i == 1) && ch == '1') {
			CleanedUp += string(ch)
		}
	}

	if len(CleanedUp) != 10 {
		return "", errors.New("invalid number")
	}
	if CleanedUp[0] < '2' || CleanedUp[3] < '2' {
		return "", errors.New("invalid number")
	}
	return CleanedUp, nil
}

func Number(phoneNumber string) (string, error) {
	return CleanUpNumber(phoneNumber)
}

func AreaCode(phoneNumber string) (string, error) {
	CleanedUp, err := CleanUpNumber(phoneNumber)
	if err != nil {
		return "", err
	}

	return CleanedUp[:3], nil
}

func Format(phoneNumber string) (string, error) {
	CleanedUp, err := CleanUpNumber(phoneNumber)
	if err != nil {
		return "", err
	}

	area, exchange, subscriber := CleanedUp[:3], CleanedUp[3:6], CleanedUp[6:]
	return fmt.Sprintf("(%s) %s-%s", area, exchange, subscriber), nil
}
