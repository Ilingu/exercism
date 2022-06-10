package thefarm

import (
	"errors"
	"fmt"
)

type ErrorType interface {
	Error() error
}

// Implement "Error" interface
type SillyNephewError struct {
	cows int
}

func (err SillyNephewError) Error() error {
	return fmt.Errorf("silly nephew, there cannot be %d cows", err.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	amount, err := weightFodder.FodderAmount()
	if err == ErrScaleMalfunction && amount >= 0 {
		amount *= 2
	} else if err == ErrScaleMalfunction && amount < 0 {
		return 0, errors.New("negative fodder")
	} else if err != nil {
		return 0, errors.New("non-scale error")
	}
	if cows == 0 {
		return 0, errors.New("division by zero")
	}
	if amount < 0 {
		return 0, errors.New("negative fodder")
	}
	if cows < 0 {
		/* Three ways:
		1. "SillyNephewError{cows: cows}.Error()" --> Only use his struct
		2. "ErrorType.Error(SillyNephewError{cows: cows})" --> Polymorphism (with interface, better for scalability)
		3. "fmt.Errorf("silly nephew, there cannot be %d cows", cows)" --> brute force method 😅
		*/
		return 0, ErrorType.Error(SillyNephewError{cows: cows})
	}

	return amount / float64(cows), nil
}
