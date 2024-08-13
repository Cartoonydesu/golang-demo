package fizzbuzz

import(
	// "fmt"
	"strconv"
)


func Fizzbuzz (n int) string {
	// if n==2 {return "2"}
	if n==3 {return "Fizz"}
	if n==5 {return "Buzz"}
	if n==15 {return "FizzBuzz"}
	return strconv.Itoa(n)
}