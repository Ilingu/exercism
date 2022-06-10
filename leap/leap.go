package leap

func IsLeapYear(input int) bool {
	return (input%4 == 0 && input%100 != 0) || (input%4 == 0 && input%400 == 0)
}
