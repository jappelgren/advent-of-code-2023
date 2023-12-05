package main

import (
	"fmt"
	"regexp"
	"math"
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
	parsedInput := CleanInput(f)

	inputMap := map[string]bool{}
	cardScore := 0

	for _, v := range parsedInput {
		tempNums := regexp.MustCompile("\\s").Split(v[0], -1)
		tempWins := regexp.MustCompile("\\s").Split(v[1], -1)
		for _, w := range tempWins {
				inputMap[w] = true
		}
		for _, w := range tempNums {
			if _, ok := inputMap[w]; ok {
				cardScore++
			}
		}
		if cardScore == 1 {
			res += cardScore
		} else {
			res += int(math.Pow(float64(2), float64(cardScore - 1)))
		}
		cardScore = 0
		inputMap = map[string]bool{}
	}
	return
}

func SecondStar(f []string) (res int) {
	parsedInput := CleanInput(f)

	inputMap := map[string]bool{}
	ticketMap := map[int]int{}
	cardScore := 0

	for i := range parsedInput {
		ticketMap[i] = 1
	}

	for i, v := range parsedInput {
		tempNums := regexp.MustCompile("\\s").Split(v[0], -1)
		tempWins := regexp.MustCompile("\\s").Split(v[1], -1)
		for _, w := range tempWins {
				inputMap[w] = true
		}
		for _, w := range tempNums {
			if _, ok := inputMap[w]; ok {
				cardScore++
				ticketMap[i + cardScore] = ticketMap[i + cardScore] + ticketMap[i]
			}
		}
		cardScore = 0
		inputMap = map[string]bool{}
	}

	for _, v := range ticketMap {
		res += v
	}
	return
}

func CleanInput(f []string) (res [][]string) {
	for _, v := range f {
		noGame := regexp.MustCompile("Card\\s{1,}[0-9]{1,}\\:\\s{1,}").ReplaceAll([]byte(v), []byte(""))
		oneSpacedSingles := regexp.MustCompile("\\s{2,}").ReplaceAll([]byte(noGame), []byte(" "))
		splitNumWins := regexp.MustCompile("\\s\\|\\s").Split(string(oneSpacedSingles), -1)
		res = append(res, splitNumWins)
	}
	return
}