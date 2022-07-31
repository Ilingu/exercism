package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) (acc int) {
	acc = initial
	for _, sc := range s {
		acc = fn(acc, sc)
	}
	return
}

func (s IntList) Foldr(fn func(int, int) int, initial int) (acc int) {
	acc = initial
	for i := len(s) - 1; i >= 0; i-- {
		acc = fn(s[i], acc)
	}
	return
}

func (s IntList) Filter(fn func(int) bool) IntList {
	out := IntList{}
	for _, sc := range s {
		if fn(sc) {
			out = append(out, sc)
		}
	}
	return out
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	out := IntList{}
	for _, sc := range s {
		out = append(out, fn(sc))
	}
	return out
}

func (s IntList) Reverse() IntList {
	out := IntList{}
	for i := len(s) - 1; i >= 0; i-- {
		out = append(out, s[i])
	}
	return out
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	out := append(IntList{}, s...)
	for _, list := range lists {
		out = append(out, list...)
	}
	return out
}
