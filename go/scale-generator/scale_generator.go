package scale

import (
	"strings"
)

type pitchesType string

const (
	sharp = "sharp"
	flat  = "flat"
)

var (
	naturalKeys = map[string]bool{"C": true, "a": true}
	sharpKeys   = map[string]bool{"G": true, "D": true, "A": true, "E": true, "B": true, "F#": true, "e": true, "b": true, "f#": true, "c#": true, "g#": true, "d#": true}
)

var ChromaticScaleSharp, ChromaticScaleFlat = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}, []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

func Scale(tonic, interval string) []string {
	indexOrder := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	keySharpIndex, AssociatedChromaticScale := tonicIndex(tonic)
	indexOrder = append(indexOrder[keySharpIndex:], indexOrder[:keySharpIndex]...)

	if interval != "" {
		for j := range indexOrder {
			if j-1 > len(interval)-1 {
				break
			}

			if j != 0 {
				CurrInterval := rune(interval[j-1])

				switch CurrInterval {
				case 'M':
					indexOrder = append(indexOrder[:j], indexOrder[j+1:]...) // skip
				case 'A':
					indexOrder = append(indexOrder[:j], indexOrder[j+2:]...) // skip 2 notes
				}
			}
		}
		indexOrder = append(indexOrder, indexOrder[0])
	}

	scale := []string{}
	for _, i := range indexOrder {
		scale = append(scale, AssociatedChromaticScale[i])
	}

	return scale
}

func sharpOrFlat(tonic string) pitchesType {
	isNatural, isSharp := naturalKeys[tonic], sharpKeys[tonic]
	if isNatural || isSharp {
		return sharp
	}
	return flat
}

func tonicIndex(tonic string) (int, []string) {
	pitches := sharpOrFlat(tonic)

	arr := ChromaticScaleFlat
	if pitches == sharp {
		arr = ChromaticScaleSharp
	}

	for i := range arr {
		if strings.EqualFold(arr[i], tonic) {
			return i, arr
		}
	}
	return 0, nil
}
