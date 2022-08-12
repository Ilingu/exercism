package allergies

import "math"

var AllergiesIdsToNames = map[int]string{1: "eggs", 2: "peanuts", 4: "shellfish", 8: "strawberries", 16: "tomatoes", 32: "chocolate", 64: "pollen", 128: "cats"}

type IntArr []int

func Allergies(allergies uint) []string {
	allergiesIds, id := make(IntArr, 1), 0
	for i := float64(0); allergiesIds.Sum() != int(allergies); i++ {
		if id > len(allergiesIds)-1 {
			allergiesIds = append(allergiesIds, 0)
		}

		allergiesIds[id] = int(math.Pow(2, i))
		if allergiesIds.Sum() > int(allergies) {
			allergiesIds[id] = int(math.Pow(2, i-1))
			id++
			i = 0
		}
	}

	allergiesNames := []string{}
	for _, allergieId := range allergiesIds {
		allergieName, ok := AllergiesIdsToNames[allergieId]
		if !ok || allergieName == "" {
			continue
		}
		allergiesNames = append(allergiesNames, allergieName)
	}

	return allergiesNames
}

func AllergicTo(allergies uint, allergen string) bool {
	allergiesNames := Allergies(allergies)
	for _, name := range allergiesNames {
		if allergen == name {
			return true
		}
	}

	return false
}

func (arr IntArr) Sum() (sum int) {
	for _, num := range arr {
		sum += num
	}
	return
}
