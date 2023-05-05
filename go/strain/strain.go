package strain

type Ints []int
type Lists [][]int
type Strings []string

// const (
// 	keep    = 0
// 	discard = 1
// )

/* USING GENERICS 😭 */
// func Handle[T any](job int, value []T, filter func(T) bool) []T {
// 	out := make([]T, 0)

// 	for _, val := range value {
// 		if (job == keep && filter(val)) || (job == discard && !filter(val)) {
// 			out = append(out, val)
// 		}
// 	}

// 	if len(out) == 0 {
// 		return nil
// 	}
// 	return out
// }

func (i Ints) Keep(filter func(int) bool) Ints {
	out := make([]int, 0)

	for _, val := range i {
		if filter(val) {
			out = append(out, val)
		}
	}

	if len(out) == 0 {
		return nil
	}
	return out
}

func (i Ints) Discard(filter func(int) bool) Ints {
	out := make([]int, 0)

	for _, val := range i {
		if !filter(val) {
			out = append(out, val)
		}
	}

	if len(out) == 0 {
		return nil
	}
	return out
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	out := make([][]int, 0)

	for _, val := range l {
		if filter(val) {
			out = append(out, val)
		}
	}

	if len(out) == 0 {
		return nil
	}
	return out
}

func (s Strings) Keep(filter func(string) bool) Strings {
	out := make([]string, 0)

	for _, val := range s {
		if filter(val) {
			out = append(out, val)
		}
	}

	if len(out) == 0 {
		return nil
	}
	return out
}
