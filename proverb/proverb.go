package proverb

import "fmt"

func Proverb(rhymes []string) []string {
	if len(rhymes) <= 0 {
		return nil
	}

	proverb := make([]string, len(rhymes)-1)
	for i := range proverb {
		proverb[i] = fmt.Sprintf("For want of a %s the %s was lost.", rhymes[i], rhymes[i+1])
	}

	proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhymes[0]))
	return proverb
}
