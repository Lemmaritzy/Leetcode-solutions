package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	num := 1000021

	isPalindrome := IsPalindrome(num)
	println(isPalindrome)
}

func IsPalindrome(x int) (status bool) {

	reverseSlice := []string{}

	if x < 0 {
		return false
	}

	stringNum := strconv.Itoa(x)
	mySlice := strings.Split(stringNum, "")

	for i := 0; i < len(mySlice); i++ {
		reverseSlice = append(reverseSlice, mySlice[len(mySlice)-i-1])
	}

	for i := 0; i < len(mySlice); i++ {
		fmt.Printf("index=%v ==> first slice=> %v || second slice=> %v \n", i, mySlice[i], reverseSlice[i])
		if mySlice[i] == reverseSlice[i] {
			status = true
		} else {
			return false
		}
	}

	return
}
