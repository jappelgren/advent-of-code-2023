package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	input := ParseFileToStringByNewLine("input.txt")
	first := FirstStar(input)
	fmt.Printf(`Day 2 First Star Result: %v%v`, first, "\r\n")

	second := SecondStar(input)
	fmt.Printf(`Day 2 Second Star Result: %v%v`, second, "\r\n")
}

func FirstStar(f []string) int {
	var res int
	maxR, maxG, maxB := 12, 13, 14

	for i, v := range f {
		gameId := i + 1
		split := regexp.MustCompile("[0-9]{2,}\\s[a-zA-Z]{3,5}").FindAllString(v, -1)

		valid := true
		for _, w := range split {
			pair := regexp.MustCompile("[0-9]{2,}|[a-zA-Z]{3,5}").FindAllString(w, -1)
			if cubes, err := strconv.Atoi(pair[0]); err == nil &&
				((pair[1] == "red" && cubes > maxR) || (pair[1] == "green" && cubes > maxG) || (pair[1] == "blue" && cubes > maxB)) {
				valid = false
				break
			}
		}

		if len(split) == 0 || valid {
			res += gameId
		}
	}

	return res
}

func SecondStar(f []string) int {
	var res int

	for _, v := range f {
		gameless := regexp.MustCompile("Game\\s[0-9]{1,}:").ReplaceAll([]byte(v), []byte(""))
		split := regexp.MustCompile("[0-9]{1,}\\s[a-zA-Z]{3,5}").FindAllString(string(gameless), -1)

		maxMap := map[string]int{"red": 0, "blue": 0, "green": 0}
		for _, w := range split {
			pair := regexp.MustCompile("[0-9]{1,}|[a-zA-Z]{3,5}").FindAllString(w, -1)
			if cubes, err := strconv.Atoi(pair[0]); err == nil && cubes > maxMap[pair[1]] {
				maxMap[pair[1]] = cubes
			}
		}

		res += maxMap["red"] * maxMap["blue"] * maxMap["green"]
	}

	return res
}
