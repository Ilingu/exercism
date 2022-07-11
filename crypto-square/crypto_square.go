package cryptosquare

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

type CryptoSquare struct {
	c           int
	r           int
	ChMissing   int
	PerfectRect bool
}

func normalizeText(str string) (nomalizedStr string) {
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)
	for _, ch := range str {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			nomalizedStr += string(ch)
		}
	}
	return
}

func NewCryptoSquare(length int) CryptoSquare {
	var c, r int
	for ic := 1; ic <= length; ic++ {
		ir := int(math.Ceil(float64(length) / float64(ic)))
		if ir*ic >= length && ic >= ir && ic-ir <= 1 {
			c, r = ic, ir
			break
		}
	}

	ChMissing := c*r - length
	return CryptoSquare{c: c, r: r, ChMissing: ChMissing, PerfectRect: ChMissing == 0}
}

// Encode Top to Down - Right to Left
func (rect CryptoSquare) EncodeTtDRtL(strToEncode string) (encodedStr string) {
	PerfectRectStr := strToEncode
	for i := 0; i < rect.ChMissing; i++ {
		PerfectRectStr += " "
	}

	RectChunks := []string{}
	for i := 0; i < rect.r; i++ {
		RectChunks = append(RectChunks, PerfectRectStr[i*rect.c:(i+1)*rect.c])
	}

	for i := 0; i < rect.c; i++ {
		for _, chunk := range RectChunks {
			encodedStr += string(chunk[i])
		}
	}
	return
}

func (rect CryptoSquare) CutInChuck(encodedStr string) (out string) {
	for i, ch := range encodedStr {
		if i%rect.r == 0 && i != 0 {
			out += " "
		}
		out += string(ch)
	}
	return
}

func Encode(pt string) string {
	pt = normalizeText(pt)
	cryptoRect := NewCryptoSquare(len(pt))

	encodedStr := cryptoRect.EncodeTtDRtL(pt)

	outEncoded := cryptoRect.CutInChuck(encodedStr)
	return outEncoded
}

func (r CryptoSquare) String() string {
	return fmt.Sprintf("columns: %d, rows: %d\nNumber of characters missing to be a perfect rect: %d\n", r.c, r.r, r.ChMissing)
}

/* bullshit
r * c >= length(message)
c >= r
c - r <= 1

---------

r >= l/c
c >= l/r

c <= 1 + r
r >= c - 1

l <= r² + r
l <= c² - c

c² - c = r² + r
c² - r² = r + c
*/
