package main

import (
	// "net/http"
	"fmt"
	// "strings"

	// "github.com/gin-gonic/gin"
)

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run()


	// hello()

	// fmt.Println(squareArea(128))

	// fmt.Println(greet("Gopher"))

	// name := "Gopher"
	// greet2(&name)
	// fmt.Println(name)
	
	list := [4]int{1, 3, 4, 2}
	r := reverse(list)
	fmt.Printf("%T %v\n", r, r)
	fmt.Printf("%v === %v\n",list, reverseSlice(list))

	// fmt.Println(sliceString("abcdef"))
	// fmt.Println(sliceString("abcdefg"))

	// test();
}

func greet(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func greet2(name *string) {
	*name = fmt.Sprintf("Hello, %s", *name)
// fmt.Println(*name)
}

func hello() {
	fmt.Println("Hello")
	fmt.Println("What is your name?: ")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Hi %s.\n", name)
}

func reverse(arr [4]int) [4]int{
	var result [4]int
	i :=len(arr)-1
	j := 0
	for i>=0 {
		result[j] = arr[i]
		i--
		j++
	} 
	return result
}

func reverseSlice(arr [4]int) []int {
	var slice []int = arr[:]
	var result []int
	for len(slice) > 0 {
		result = append(result, slice[len(slice)-1])
		slice = slice[:len(slice)-1]
		// break
	}
	return result
	// return []int{arr[3], arr[2], arr[1], arr[0]}
}

func sliceString(s string) []string{
	var result []string
	// // i, j := 0,1
	// // i:= 0
	// sArr := strings.SplitAfter(s, "")
	// fmt.Print(sArr)
	// i:= 0
	// for i< len(sArr)/2 {
	// 	// fmt.Println(i, i+1)
	// 	fmt.Println(sArr[i:i+2])
	// 	result = append(result, sArr[i:i+2])
	// 	// j++
	// 	i = i+2
	// }
	// return result
	for s+="*"; len(s)>1; s = s[2:] {
		result = append(result, s[:2])
	}
	return result
}

func squareArea(a float64) float64 {
	return a*a
}

func test() {
	p:=10
	var i = &p
	var j = &i
	fmt.Printf("%v %p %v\n", &i, i, *i)
	fmt.Printf("j => %v %p %v %v\n", &j, j, *j, **j)
	test2(i, j)
}

func test2(i *int, j **int) {
	fmt.Printf("%v %p %v\n", &i, i, *i)
	fmt.Printf("%v %p %v\n", &j, j, *j)
	fmt.Printf("j => %v %p %v %v\n", &j, j, *j, **j)
}

