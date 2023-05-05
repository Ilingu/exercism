package pythagorean

type Triplet [3]int

func (t Triplet) perimeter() int {
	return t[0] + t[1] + t[2]
}

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	var IterableRange []int
	for i := min; i <= max; i++ {
		IterableRange = append(IterableRange, i)
	}

	var ValidTriplets []Triplet
	for i, a := range IterableRange {
		for j, b := range IterableRange[i+1:] {
			a2b2 := a*a + b*b
			for _, c := range IterableRange[j+1:] {
				c2 := c * c
				if a2b2 == c2 {
					ValidTriplets = append(ValidTriplets, Triplet{a, b, c})
				}
				if c2 >= a2b2 {
					break
				}
			}
		}
	}

	return ValidTriplets
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	PyRanges := Range(1, p)

	var ValidTriplets []Triplet
	for _, triplet := range PyRanges {
		if triplet.perimeter() == p {
			ValidTriplets = append(ValidTriplets, triplet)
		}
	}

	return ValidTriplets
}
