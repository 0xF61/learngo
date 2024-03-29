package chance

import "math/rand"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return 1 + rand.Intn(20)
	// panic("Please implement the RollADie function")
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return float64(rand.Intn(12)) + rand.Float64()
	// panic("Please implement the GenerateWandEnergy function")
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}

	swap := func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	}
	rand.Shuffle(len(animals), swap)

	return animals
	// panic("Please implement the ShuffleAnimals function")
}
