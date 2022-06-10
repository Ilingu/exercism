package robotname

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// Define the Robot type here.
type Robot struct {
	name string
}

var Letters = [26]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
var NameSeen = map[string]bool{"": true}

func RandNum(limit int) int64 {
	/* My First Method before discovering "rand.Int" 😭 --> DO NOT USE, bad performance */
	// RandomCrypto, _ := rand.Prime(rand.Reader, 128)
	// res := big.NewInt(0).Div(RandomCrypto, big.NewInt(int64(limit*2)))
	// res = big.NewInt(0).Add(res, RandomCrypto)
	// res = big.NewInt(0).Mod(res, big.NewInt(int64(limit)))
	// return res.Int64()
	res, _ := rand.Int(rand.Reader, big.NewInt(int64(limit)))
	return res.Int64()
}

func GenerateName() string {
	limit := len(Letters)

	LetterA, LetterB := Letters[RandNum(limit)], Letters[RandNum(limit)]
	digitA, digitB, digitC := RandNum(10), RandNum(10), RandNum(10)

	return fmt.Sprintf("%s%s%d%d%d", string(LetterA), string(LetterB), digitA, digitB, digitC)
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	name := ""
	for NameSeen[name] {
		name = GenerateName()
	}
	NameSeen[name] = true

	r.name = name
	return name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}
