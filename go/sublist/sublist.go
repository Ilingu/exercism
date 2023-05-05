package sublist

type Relation string

const (
	equal     = "equal"
	sublist   = "sublist"
	superlist = "superlist"
	unequal   = "unequal"
)

func IsSublist(smaller, bigger []int) bool {
	if len(smaller) == 0 {
		return true
	}
	if len(smaller) > len(bigger) {
		return false
	}

	for i := range bigger[:len(bigger)-len(smaller)+1] {
		for j := range smaller {
			if smaller[j] != bigger[i+j] {
				break
			}
			if j == len(smaller)-1 {
				return true
			}
		}
	}

	return false
}

func Sublist(l1, l2 []int) Relation {
	if len(l1) == len(l2) && IsSublist(l1, l2) {
		return equal
	}

	if IsSublist(l1, l2) {
		return sublist
	}

	// Not Sublist --> Check Superlist
	if IsSublist(l2, l1) {
		return superlist
	}

	return unequal // If nor equal nor sublist nor superlist, it have to be unequal
}
