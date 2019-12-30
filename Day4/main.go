package main

import (
	"errors"
	"fmt"
)

const LOWER_BOUND = 168630
const HIGHER_BOUND = 718098

type rule func(int) (bool, error)
type validatable func(...rule) (bool, error)

func main() {
	var ctr int

	for i := LOWER_BOUND; i <= HIGHER_BOUND; i++ {
		ok, _ := PasswordValidate(i)

		if ok {
			ctr++
		}
	}

	fmt.Println(ctr)
}

func PasswordValidate(input int) (bool, error) {
	validator := Validator(input)

	return validator(
		outOfBounds,
		nonDecreasingDigits,
		onlyRepeatsTwice,
	)
}

func Validator(input int) validatable {
	return func(rules ...rule) (bool, error) {
		for _, r := range rules {
			ok, msg := r(input)
			if ok == false {
				return false, msg
			}
		}

		return true, nil
	}
}

func outOfBounds(input int) (bool, error) {
	if input < LOWER_BOUND || input > HIGHER_BOUND {
		return false, errors.New("The input provided is out of bounds")
	}

	return true, nil
}

func duplicate(input int) (bool, error) {
	temp := input % 10

	for i := input / 10; i > 0; i /= 10 {
		if temp == i%10 {
			return true, nil
		}

		temp = i % 10
	}

	return false, errors.New("The input provided does not contain any duplicate pairs")
}

func nonDecreasingDigits(input int) (bool, error) {
	temp := input % 10

	for i := input / 10; i > 0; i /= 10 {
		if temp < i%10 {
			return false, errors.New("The input is not in a non-decreasing format")
		}

		temp = i % 10
	}

	return true, nil
}

func onlyRepeatsTwice(input int) (bool, error) {
	var container Container

	for i := input; i > 0; i /= 10 {
		if container.last() != i%10 && len(container) == 2 {
			return true, nil
		}

		container.store(i % 10)
	}

	if container.valid() {
		return true, nil
	}

	return false, errors.New("The input does not contain 1 pair of same number")
}
