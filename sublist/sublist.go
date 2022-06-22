package sublist

type Relation string

const (
	equal     = "equal"
	sublist   = "sublist"
	superlist = "superlist"
	unequal   = "unequal"
)

func IsEqual(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}
	// No matter whether I pick len(l2) or len(l1) since they have the same length
	for i := 0; i < len(l2); i++ {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}

func IsSublist(smaller, bigger []int) bool {
	if len(smaller) >= len(bigger) {
		return false
	}

	SublistFound := true

	var OrderedFound, LastIndexEq int
	for i, BigVal := range bigger {
		for j, SmVal := range smaller {
			SublistFound = false

			if SmVal != BigVal {
				continue // If no equality, we pass
			}

			if j == 0 {
				OrderedFound = 1 // New Potential Sublist Detected
			}

			// If true, that means there is a no gap between two element of the sublist, they follow each other
			if LastIndexEq+1 == i {
				OrderedFound++
				if OrderedFound == len(smaller)-1 {
					return true // Sublist found
				}
			} else {
				OrderedFound = 0
			}

			LastIndexEq = i // Keeping track of where was the last time a equality between L1 and L2 occur
			break
		}
		// If no other natural possibility
		if i > len(bigger)-(len(bigger)-len(smaller)) && OrderedFound == 0 {
			break
		}
	}

	return SublistFound
}

func Sublist(l1, l2 []int) Relation {
	if IsEqual(l1, l2) {
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
