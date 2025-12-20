package lasagna

func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}

	return len(layers) * timePerLayer
}

func Quantities(layers []string) (int, float64) {
	noodlesQty := 50
	sauceQty := 0.2

	totalNoodlesQty := 0
	totalSouceQty := 0.0

	for _, v := range layers {
		if v == "noodles" {
			totalNoodlesQty += noodlesQty
		} else if v == "sauce" {
			totalSouceQty += sauceQty
		}
	}

	return totalNoodlesQty, totalSouceQty
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendIngredients, recipeIngredients []string) {
	lastItem := friendIngredients[len(friendIngredients)-1]

	recipeIngredients[len(recipeIngredients)-1] = lastItem
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	newRecipe := make([]float64, len(quantities))
	multiplier := float64(portions) / 2.0

	for i, portion := range quantities {
		newRecipe[i] = portion * float64(multiplier)
	}

	return newRecipe
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
