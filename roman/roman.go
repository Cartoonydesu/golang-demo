package roman

import (
	"strings"
)

// func RomanCalculator(num string) int {
// 	var result int
// 	var numSplit = strings.Split(num, "")
// 	var convertNum []int
// 	for _, element := range numSplit {
// 		switch element {
// 		case "I": convertNum = append(convertNum, 1)
// 		case "V": convertNum = append(convertNum, 5)
// 		case "X": convertNum = append(convertNum, 10)
// 		case "L": convertNum = append(convertNum, 50)
// 		case "C": convertNum = append(convertNum, 100)
// 		case "D": convertNum = append(convertNum, 500)
// 		case "M": convertNum = append(convertNum, 1000)
// 		}
// 	}
// 	count0, count1 := 0,1
// 	for count0 < len(convertNum) {
// 		currentNum := convertNum[count0]
// 		if count1 >= len(convertNum) {
// 			result += currentNum
// 			break 
// 		} else {
// 			nextNum := convertNum[count1]
// 			if(nextNum > currentNum) {
// 				result -= currentNum
// 			} else {
// 				result += currentNum
// 			}
// 		}
// 		count0++
// 		count1++
// 	}

// 	return result
// }

func RomanCalculator(num string) int {
	var result int
	var numSplit = strings.Split(num, "")
	var first = RomanDictionary(numSplit[0])
	if (len(numSplit) > 1) {
		for i:=1; i<len(numSplit); i++ {
			var second = RomanDictionary(numSplit[i])
			if (first < second) {
				result -= first
			} else {
				result += first
			}
			first = second
		}
	} 
	result += first
	// var second = RomanDictionary(numSplit[1])
	return result
}

func RomanDictionary(numString string) int {
	switch numString {
		case "I": return 1
		case "V": return 5
		case "X": return 10
		case "L": return 50
		case "C": return 100
		case "D": return 500
		case "M": return 1000
	}
	return 0
}