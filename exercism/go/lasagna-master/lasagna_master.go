package lasagna

// Calculate PreparationTime given preptime for each layer.
func PreparationTime(layers []string, layerPrepTime int) int {
	layerPrepTime = map[bool]int{true: 2, false: layerPrepTime}[layerPrepTime <= 0]

	return layerPrepTime * len(layers)
}

// Calculate required noodle & sauce quantities.
func Quantities(layers []string) (int, float64) {
	var noodlePerLayer int = 50
	var saucePerLayer float64 = .2

	var requiredNoodles int
	var requiredSauces float64

	for _, v := range layers {
		switch v {
		case "noodles":
			requiredNoodles += noodlePerLayer
		case "sauce":
			requiredSauces += saucePerLayer
		}
	}

	return requiredNoodles, requiredSauces
}

// Copy the secret(last) ingredient from friendList to myList
func AddSecretIngredient(friendList, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// Scale the quantities according to portion
func ScaleRecipe(quantities []float64, portion int) []float64 {
	var scaledQuantities []float64
	scaledQuantities = append(scaledQuantities, quantities...)

	for i, quantity := range quantities {
		scaledQuantities[i] = (quantity * float64(portion)) / 2
	}

	return scaledQuantities
}
