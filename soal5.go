package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var str string
	fmt.Println("input kata :")
	fmt.Scan(&str)
	result := isPalindrome(str)

	fmt.Println(result)
}
