package main

import "fmt"

func reverseString(s string) string {
	strings := []rune(s)
	for i, j := 0, len(strings)-1; i < j; i, j = i+1, j-1 {
		strings[i], strings[j] = strings[j], strings[i]
	}
	return string(strings)
}

func main() {
	var str string
	fmt.Println("input kata :")
	fmt.Scan(&str)
	fmt.Println(reverseString(str))
}
