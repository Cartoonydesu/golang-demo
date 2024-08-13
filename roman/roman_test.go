package roman

import (
	"testing"
)

func TestRomanNumberOne(t *testing.T) {
	given := "I"
	want := 1
	result := RomanCalculator(given)
	if want != result {
		t.Errorf("given %v wants %v but %v\n", given, want, result)
	}
}

func TestRomanNumberTwo(t *testing.T) {
	given := "II"
	want := 2
	result := RomanCalculator(given)
	if want != result {
		t.Errorf("given %v wants %v but %v\n", given, want, result)
	}
}

func TestRomanNumberThree(t *testing.T) {
	given := "III"
	want := 3
	result := RomanCalculator(given)
	if want != result {
		t.Errorf("given %v wants %v but %v\n", given, want, result)
	}
}

//4

func TestRomanNumberFour(t *testing.T) {
	given := "IV"
	want := 4
	result := RomanCalculator(given)
	if want != result {
		t.Errorf("given %v wants %v but %v\n", given, want, result)
	}
}

func TestRomanNumberFive(t *testing.T) {
	given := "V"
	want := 5
	result := RomanCalculator(given)
	if want != result {
		t.Errorf("given %v wants %v but %v\n", given, want, result)
	}
}

func TestRomanNumberSix(t *testing.T) {
	given := "VI"
	want := 6
	result := RomanCalculator(given)
	if want != result {
		t.Errorf("given %v wants %v but %v\n", given, want, result)
	}
}

// 7 8 9

func TestRomanNumberNine(t *testing.T) {
	given := "IX"
	want := 9
	result := RomanCalculator(given)
	if want != result {
		t.Errorf("given %v wants %v but %v\n", given, want, result)
	}
}

func TestRomanNumberTen(t *testing.T) {
	given := "X"
	want := 10
	result := RomanCalculator(given)
	if want != result {
		t.Errorf("given %v wants %v but %v\n", given, want, result)
	}
}

func TestRomanNumberSixtyFour(t *testing.T) {
	given := "LXIV"
	want := 64
	result := RomanCalculator(given)
	if want != result {
		t.Errorf("given %v wants %v but %v\n", given, want, result)
	}
}

func TestRomanNumberOneHundredFiftySix(t *testing.T) {
	given := "CLVI"
	want := 156
	result := RomanCalculator(given)
	if want != result {
		t.Errorf("given %v wants %v but %v\n", given, want, result)
	}
}

func TestRomanNumberNineHundredNinetyNine(t *testing.T) {
	given := "CMXCIX"
	want := 999
	result := RomanCalculator(given)
	if want != result {
		t.Errorf("given %v wants %v but %v\n", given, want, result)
	}
}
