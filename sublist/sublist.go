package sublist

import (
	"fmt"
	"strings"
)

type Relation string

func Sublist(l1, l2 []int) Relation {
	cutArray := func(str string) string {
		return strings.ReplaceAll(strings.ReplaceAll(str, "[", ""), "]", "")
	}
	l1Sig, l2Sig := cutArray(fmt.Sprint(l1)), cutArray(fmt.Sprint(l2))

	if l1Sig == l2Sig {
		return "equal"
	} else if strings.Contains(l2Sig, l1Sig) {
		return "sublist"
	} else if strings.Contains(l1Sig, l2Sig) {
		return "superlist"
	}
	return "unequal"
}
