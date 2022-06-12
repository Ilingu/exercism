package protein

import (
	"fmt"
)

type CustomError struct {
	message string
}

func NewError(msg string) *CustomError {
	return &CustomError{message: msg}
}

func (e *CustomError) Error() string {
	return e.message
}

var ErrStop = NewError("error stop")
var ErrInvalidBase = NewError("error invalid base")

var FromCodons = map[string]string{"AUG": "Methionine", "UUU": "Phenylalanine", "UUC": "Phenylalanine", "UUA": "Leucine", "UUG": "Leucine", "UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine", "UAU": "Tyrosine", "UAC": "Tyrosine", "UGU": "Cysteine", "UGC": "Cysteine", "UGG": "Tryptophan", "UAA": "STOP", "UAG": "STOP", "UGA": "STOP"}

// func FromProtein() map[string][]string {
// 	fromProtein := make(map[string][]string)
// 	for codon, protein := range FromCodons {
// 		fromProtein[protein] = append(fromProtein[protein], codon)
// 	}
// 	return fromProtein
// }

func FromRNA(rna string) ([]string, error) {
	if len(rna)%3 != 0 {
		return nil, ErrInvalidBase
	}

	Proteins := []string{}
	CodonsNum := len(rna) / 3
	for i := 0; i < CodonsNum; i++ {
		codon := rna[i*3 : i*3+3]
		protein, err := FromCodon(codon)
		if err == ErrStop {
			break
		}
		if err == ErrInvalidBase {
			return nil, err
		}
		Proteins = append(Proteins, protein)
	}

	return Proteins, nil
}

func FromCodon(codon string) (string, error) {
	protein, ok := FromCodons[codon]
	if !ok {
		return "", ErrInvalidBase
	}
	fmt.Println(codon, protein)
	if protein == "STOP" {
		return "", ErrStop
	}
	return protein, nil
}
