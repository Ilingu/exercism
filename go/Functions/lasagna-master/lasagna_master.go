package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
	// return len(layers) * (timePerLayer | 2)
	if timePerLayer <= 0 {
		return len(layers) * 2
	}
	return len(layers) * timePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, v := range layers {
		if v == "noodles" {
			noodles += 50
		}
		if v == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, amounts int) []float64 {
	CopyQuantitites := make([]float64, len(quantities))
	copy(CopyQuantitites, quantities)

	for i := 0; i < len(CopyQuantitites); i++ {
		CopyQuantitites[i] = float64(amounts) * CopyQuantitites[i] / 2
	}
	return CopyQuantitites
}
