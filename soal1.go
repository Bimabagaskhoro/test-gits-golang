package main

import "fmt"

func main() {
	var i int
	fmt.Println("input angka :")
	fmt.Scan(&i)

	if i%3 == 0 && i%5 == 0 {
		fmt.Println("Hello World")
	} else if i%5 == 0 {
		fmt.Println("Hello")
	} else {
		fmt.Println("World")
	}

}
