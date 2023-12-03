package main

import ("fmt")

func main (){
	input := ParseFileToStringByNewLine("example.txt")
	fmt.Println("Welcome to day 3")

	first := FirstStar(input)
	fmt.Printf(`First Star result: %v%v`, first, "\r\n")
	
	second := SecondStar(input)
	fmt.Printf(`Second Star result: %v%v`, second, "\r\n")


}

func FirstStar (f []string) int {
	return 0
}

func SecondStar (f []string) int {
	return 0
}