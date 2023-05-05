package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(NumberBox.Number(nb)))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	TFN, ok := fnb.(FancyNumber)
	NumberOutput, err := strconv.Atoi(TFN.Value())
	if !ok || err != nil {
		return 0
	}
	return NumberOutput
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	boxNumber := ExtractFancyNumber(fnb)
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(boxNumber))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch t := i.(type) {
	case int:
		return DescribeNumber(float64(t))
	case float64:
		return DescribeNumber(t)
	case NumberBox:
		return DescribeNumberBox(t)
	case FancyNumberBox:
		return DescribeFancyNumberBox(t)
	default:
		return "Return to sender"
	}
}
