package main

import (
	"fmt"
)

func main() {
	input := ParseFileToStringByNewLine("input.txt")
	fmt.Println("Welcome to day 4")

	first := FirstStar(input)
	fmt.Printf(`First Star result: %v%v`, first, "\r\n")

	second := SecondStar(input)
	fmt.Printf(`Second Star result: %v%v`, second, "\r\n")
}

func FirstStar(f []string) (res int) {
	return
}

func SecondStar(f []string) (res int) {
	return
}