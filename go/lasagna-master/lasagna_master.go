package lasagna_master

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgTime int) int {
	if avgTime == 0 {
		return len(layers) * 2
	}
	return len(layers) * avgTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		}
		if layers[i] == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsRecipe, myRecipe []string) {
	myRecipe[len(myRecipe)-1] = friendsRecipe[len(friendsRecipe)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var scales = make([]float64, len(quantities))
	for i := 0; i < len(quantities); i++ {
		scales[i] = quantities[i] * float64(portions) / 2
	}
	return scales
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
