package series

func All(n int, s string) []string {
	out := make([]string, 0)
	for i := 0; i <= len(s)-n; i++ {
		out = append(out, s[i:i+n])
	}
	return out
}

func UnsafeFirst(n int, s string) (out string) {
	first, ok := First(n, s)

	if !ok {
		panic(nil)
	}

	return first
}

func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}
