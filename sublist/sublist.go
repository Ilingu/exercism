package sublist

type Relation string

const (
	equal     = "equal"
	sublist   = "sublist"
	superlist = "superlist"
	unequal   = "unequal"
)

func IsSublist(smaller, bigger []int, inversed ...bool) Relation {
	SublistFound := true
	var OrderedFound, LastIndexEq int
	for i, SmVal := range smaller {
		for j, BigVal := range bigger {
			SublistFound = false

			eq := SmVal == BigVal
			if !eq {
				continue // If no equality, we pass
			}

			if i == 0 {
				OrderedFound = 1 // New Potential Sublist Detected
			}

			// If true, that means there is a no gap between two element of the sublist, they follow each other
			if LastIndexEq+1 == j {
				OrderedFound++
				if OrderedFound == len(smaller)-1 {
					SublistFound = true
				}
			} else {
				OrderedFound = 0
			}

			LastIndexEq = j // Keeping track of where was the last time a equality between L1 and L2 occur
			break
		}

		if SublistFound {
			break
		}
	}

	if len(smaller) == len(bigger) && SublistFound {
		return equal
	}
	if SublistFound {
		if inversed != nil {
			return superlist
		}
		return sublist
	}
	return unequal
}

func Sublist(l1, l2 []int) Relation {
	// Check Smaller, to pass the args in the right order to IsSublist
	if len(l1) > len(l2) {
		return IsSublist(l2, l1, true)
	}
	return IsSublist(l1, l2)
}
