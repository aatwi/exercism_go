package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodleCount := 0
	sauceCount := 0
	for index := range layers {
		layer := layers[index]
		if layer == "noodles" {
			noodleCount++
		} else if layer == "sauce" {
			sauceCount++
		}
	}
	return noodleCount * 50, float64(sauceCount) * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(list []string, list2 []string) {
	secret := list[len(list)-1]
	list2[len(list2)-1] = secret
}

// TODO: define the 'ScaleRecipe()' function

func ScaleRecipe(list []float64, portions int) []float64 {
	result := make([]float64, len(list))
	for i := range result {
		result[i] = list[i] * float64(portions) / 2
	}
	return result
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
