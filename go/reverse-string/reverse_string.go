package reverse

func Reverse(input string) string {
	chars := []rune(input)
	var out []rune
	for i := len(chars) - 1; i >= 0; i-- {
		out = append(out, chars[i])
	}
	return string(out)
}
