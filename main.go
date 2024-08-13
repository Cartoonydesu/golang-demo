package main

import(
	"fmt"
)

func main() {
	fmt.Println(isEven(2))
	fmt.Println(isOdd(2))
	a,b := 1,2
	fmt.Println(add(a,b))
}

func add [T int | float64] (a,b T) T {
	return a+b
}


func isEven(n int) bool {
	return n%2 == 0
	//return n&1 == 0
}

func isOdd(n int) bool {
	return n%2 != 0
	//return n&1 == 1
}
