package ocr

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

var recognizeDigit int
var hashToDigit = map[string]string{
	"3f2a286c918fa358a872878eae76e6b8391e142270fac433bef4a89e457f780b": "0", "d8dc3f586e9c7a1d74cbeff002155fc1a2a7c101e6b1b6678b97dd96e447881a": "1", "15e3b2b08f94141ee92ec84f3b35016c8af8da9f4aded96aa9e377920d0195a3": "2", "d74f28ee9f3c7dc55ea37ba588ece49d390fb48195497468ea30e81ade43557d": "3", "4975dc94d6e6ac8165609110b54c5168085730dd4914cfe7e1aa0de0500860c9": "4", "7b628e5ba5db3202e65a1853b5b861ce998f473e7b271c78ae0fd01b24f1c26f": "5", "a9429e50feb175dfcea9621f0736ce1030a2a9426706cef3138ed7f91fac37bf": "6", "7cfa5aa76d72c97c1d3b152248ee6d43ad538ec793959dc2e8f1143453e83ba7": "7", "6afcbe914386cfa27958b4985bd1671e22eaa30d9055bef3dd18009aa673709e": "8", "137327872af6095e8db476e90549a37aebf87332bea56514ca9d0e1e3e481c7c": "9",
}

type FlatDigit []string
type Lines []FlatDigit

func Recognize(digits string) []string {
	digits = strings.TrimPrefix(digits, "\n")
	digitsTableRows := strings.Split(digits, "\n")

	lines := make(Lines, len(digitsTableRows)/4)
	for i, row := range digitsTableRows {
		if (i+1)%4 == 0 {
			continue
		}

		lineIndex := i / 4
		if len(lines[lineIndex]) == 0 {
			lines[lineIndex] = append(lines[lineIndex], make([]string, len(row)/3)...)
		}

		for j, digitBody := range row {
			digitIndex := j / 3
			// if len(digitIndex) > 9 {}
			lines[lineIndex][digitIndex] += string(digitBody)
		}
	}

	linesDigits := make([]string, len(digitsTableRows)/4)
	for i, line := range lines {
		for _, digits := range line {
			hashedDigit := hash(digits)

			realDigit, ok := hashToDigit[hashedDigit]
			if !ok {
				linesDigits[i] += "?"
				continue
			}

			linesDigits[i] += realDigit
		}
	}

	return linesDigits
}

func hash(str string) string {
	ByteHash := sha256.Sum256([]byte(str))
	HashedPassword := fmt.Sprintf("%x", ByteHash[:])
	return HashedPassword
}
