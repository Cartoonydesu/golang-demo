package fizzbuzz

import(
	"testing"
)

func TestFizzBuzzOne(t *testing.T) {
	given := 1
	want := "1"
	result := Fizzbuzz(given)
	if want != result {
		t.Errorf("given %d wants %q but %q\n", given, want, result)
	}
}

func TestFizzBuzzTwo(t *testing.T) {
	given := 2
	want := "2"
	result := Fizzbuzz(given)
	if want != result {
		t.Errorf("given %d wants %q but %q\n", given, want, result)
	}
}

func TestFizzBuzzThree(t *testing.T) {
	given := 3
	want := "Fizz"
	result := Fizzbuzz(given)
	if want != result {
		t.Errorf("given %d wants %q but %q\n", given, want, result)
	}
}

func TestFizzBuzzFive(t *testing.T) {
	given := 5
	want := "Buzz"
	result := Fizzbuzz(given)
	if want != result {
		t.Errorf("given %d wants %q but %q\n", given, want, result)
	}
}

func TestFizzBuzzFifteen(t *testing.T) {
	given := 15
	want := "FizzBuzz"
	result := Fizzbuzz(given)
	if want != result {
		t.Errorf("given %d wants %q but %q\n", given, want, result)
	}
}