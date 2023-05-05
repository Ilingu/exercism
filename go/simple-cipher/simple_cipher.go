package cipher

import (
	"strings"
	"unicode"
)

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.
<<<<<<< HEAD

=======
>>>>>>> exercise
type shift struct {
	distance int
}
type vigenere struct {
	key string
}

const abc = "abcdefghijklmnopqrstuvwxyz"

func NewCaesar() Cipher {
	return shift{}
}
<<<<<<< HEAD

=======
>>>>>>> exercise
func NewShift(distance int) Cipher {
	if distance >= 26 || distance <= -26 || distance == 0 {
		return nil
	}
	return shift{distance: distance}
}
<<<<<<< HEAD

=======
>>>>>>> exercise
func (c shift) Encode(input string) (encodedStr string) {
	for _, ch := range input {
		if unicode.IsLetter(ch) {
			indexAbc := strings.Index(abc, string(unicode.ToLower(ch)))
			CipheredIndex := ((indexAbc + 3) + 26) % 26
			if c.distance != 0 {
				CipheredIndex = ((indexAbc + c.distance) + 26) % 26
			}
<<<<<<< HEAD

=======
>>>>>>> exercise
			encodedStr += string(abc[CipheredIndex])
		}
	}
	return
}
<<<<<<< HEAD

=======
>>>>>>> exercise
func (c shift) Decode(input string) (decodedStr string) {
	for _, ch := range input {
		if unicode.IsLetter(ch) {
			indexAbc := strings.Index(abc, string(unicode.ToLower(ch)))
			UncipheredIndex := ((indexAbc - 3) + 26) % 26
			if c.distance != 0 {
				UncipheredIndex = ((indexAbc - c.distance) + 26) % 26
			}
<<<<<<< HEAD

=======
>>>>>>> exercise
			decodedStr += string(abc[UncipheredIndex])
		}
	}
	return
}
<<<<<<< HEAD

=======
>>>>>>> exercise
func NewVigenere(key string) Cipher {
	var ACounter int
	for _, keyCh := range key {
		if keyCh == 'a' {
			ACounter++
		}
		if !unicode.IsLetter(keyCh) || !unicode.IsLower(keyCh) {
			return nil
		}
	}
	if len(key) == 0 || ACounter == len(key) {
		return nil
	}
	return vigenere{key: key}
}
<<<<<<< HEAD

=======
>>>>>>> exercise
func NormalizeString(str string) (normalizedInput string) {
	for _, ch := range str {
		if unicode.IsLetter(ch) {
			normalizedInput += string(ch)
		}
	}
	return
}
<<<<<<< HEAD

func (v vigenere) Encode(input string) (vigStr string) {
	normalizedInput := NormalizeString(input)

	for i, ch := range normalizedInput {
		AssociatedKey := string(v.key[i%len(v.key)])
		indexAbc, distanceKey := strings.Index(abc, string(unicode.ToLower(ch))), strings.Index(abc, AssociatedKey)

=======
func (v vigenere) Encode(input string) (vigStr string) {
	normalizedInput := NormalizeString(input)
	for i, ch := range normalizedInput {
		AssociatedKey := string(v.key[i%len(v.key)])
		indexAbc, distanceKey := strings.Index(abc, string(unicode.ToLower(ch))), strings.Index(abc, AssociatedKey)
>>>>>>> exercise
		CipheredIndex := ((indexAbc + distanceKey) + 26) % 26
		vigStr += string(abc[CipheredIndex])
	}
	return
}
<<<<<<< HEAD

func (v vigenere) Decode(input string) (unVigStr string) {
	normalizedInput := NormalizeString(input)

	for i, ch := range normalizedInput {
		AssociatedKey := string(v.key[i%len(v.key)])
		indexAbc, distanceKey := strings.Index(abc, string(unicode.ToLower(ch))), strings.Index(abc, AssociatedKey)

=======
func (v vigenere) Decode(input string) (unVigStr string) {
	normalizedInput := NormalizeString(input)
	for i, ch := range normalizedInput {
		AssociatedKey := string(v.key[i%len(v.key)])
		indexAbc, distanceKey := strings.Index(abc, string(unicode.ToLower(ch))), strings.Index(abc, AssociatedKey)
>>>>>>> exercise
		UncipheredIndex := ((indexAbc - distanceKey) + 26) % 26
		unVigStr += string(abc[UncipheredIndex])
	}
	return
}
