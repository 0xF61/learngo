package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, numberCows int) (float64, error) {
	fa, err := fc.FodderAmount(numberCows)
	if err != nil {
		return 0, err
	}

	ff, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return ((fa / float64(numberCows)) * ff), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, numberCows int) (float64, error) {
	if numberCows < 1 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(fc, numberCows)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error {
	if cows > 0 {
		return nil
	}
	s_fmt := fmt.Sprintf("%v cows are invalid: ", cows)

	if cows == 0 {
		return errors.New(s_fmt + "no cows don't need food")
	}

	return errors.New(s_fmt + "there are no negative cows")
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
