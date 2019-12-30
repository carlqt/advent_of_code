package main

import "testing"

func TestValidator(t *testing.T) {
	// when there is duplicate
	input := 112346
	validator := Validator(input)

	rule1 := func(x int) (bool, error) {
		return false, nil
	}

	rule2 := func(x int) (bool, error) {
		return true, nil
	}

	actual, _ := validator(rule1, rule2)
	expected := false
	// Expect to immediately return false

	if actual != expected {
		t.Errorf("Validator should be %t", expected)
	}
}

func TestOutOfBound(t *testing.T) {
	input := 168629

	actual, _ := outOfBounds(input)
	expected := false

	if actual != expected {
		t.Errorf("%d should be %t when input is out of range", input, expected)
	}
}

func TestDuplicate(t *testing.T) {
	input := 1233456

	actual, _ := duplicate(input)
	expected := true

	if actual != expected {
		t.Errorf("%d should be %t when input has an adjacent duplicate", input, expected)
	}

	// when input has no duplicate
	input = 1234356

	actual, _ = duplicate(input)
	expected = false

	if actual != expected {
		t.Errorf("%d should be %t when input has no adjacent duplicates", input, expected)
	}
}

func TestNonDecreasingDigits(t *testing.T) {
	input := 223450
	actual, _ := nonDecreasingDigits(input)
	expected := false

	if actual != expected {
		t.Errorf("%d should be %t when input has a decreasing pair", input, expected)
	}
}

func TestOnlyRepeatsTwice(t *testing.T) {
	input := 123444

	actual, _ := onlyRepeatsTwice(input)
	expected := false

	if actual != expected {
		t.Errorf("%d should be %t when input repeats more than twice", input, expected)
	}

	// when input has no duplicate
	input = 111122

	actual, _ = onlyRepeatsTwice(input)
	expected = true

	if actual != expected {
		t.Errorf("%d should be %t when input has an adjacent pair", input, expected)
	}
}
