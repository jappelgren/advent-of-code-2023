package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Welcome to day 7")
	f := ParseFileToStringByNewLine("input.txt")

	first := FirstStar(f, 1)
	fmt.Printf("First Star Result: %v%v", first, "\r\n")

	second := SecondStar(f)
	fmt.Printf("Second Star Result: %v%v", second, "\r\n")
}

func FirstStar(f []string, ver int) (res int) {
	handMap := map[[5]int]int{}
	handRank := [][5]int{}
	for _, v := range f {
		handWager := regexp.MustCompile("\\s").Split(v, -1)
		handSlice := [5]int{}

		for i, w := range handWager[0] {
			handSlice[i] = CardToValue(w, ver)
		}

		if num, err := strconv.Atoi(handWager[1]); err == nil {
			handMap[handSlice] = num
		}
		handRank = append(handRank, handSlice)
	}
	
	SortHands(handRank)
	
	for i, v := range handRank {
		if res == 0 {
			res = handMap[v] * (len(handRank)-i)
		} else {
			res += handMap[v] * (len(handRank)-i)
		}
	}
	return
}

func SecondStar(f []string) int {
	return FirstStar(f, 2)
}

func CardToValue(c rune, ver int) int {
	jVal := 11
	if ver == 2 {
		jVal = 1
	}

	cardValMap := map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': jVal,
		'T': 10,
	}

	if v, ok := cardValMap[c]; ok {
		return v
	}

	if num, err := strconv.Atoi(string(c)); err == nil {
		return num
	}

	return 0
}

func WhichWins(hOne [5]int, hTwo [5]int) int {
	hOneScore := ScoreHand(hOne)
	hTwoScore := ScoreHand(hTwo)

	if hOneScore > hTwoScore {
		return 0
	} else if hOneScore < hTwoScore {
		return 1
	} else {
		res := -1
		for i := 0; i < 6; i++ {
			if hOne[i] == hTwo[i] {
				continue
			} else if hOne[i] > hTwo[i] {
				res = 0
				break
			} else {
				res = 1
				break
			}
		}
		return res
	}
}

func ScoreHand(h [5]int) int {
	handMap := map[int]int{}
	for _, v := range h {
		if w, ok := handMap[v]; ok {
			handMap[v] = w + 1
		} else {
			handMap[v] = 1
		}
	}

	if joker, ok := handMap[1]; ok {
			count := 0
			key := 0
			for k, v := range handMap {
				if v >= count && k > 1 {
					key = k
					count = v
				}
			}
			handMap[key] += joker
			delete(handMap, 1)
	}

	switch len(handMap) {
	case 1:
		return 6
	case 2:
		res := 4
		for _, v := range handMap {
			if v == 4 {
				res = 5
				break
			}
		}
		return res
	case 3:
		res := 2
		for _, v := range handMap {
			if v == 3 {
				res = 3
				break
			}
		}
		return res
	case 4:
		return 1
	default:
		return 0
	}
}

func SortHands(handRank [][5]int) {
	leftWins := 0
	i := 0
	j := 1

	for leftWins < len(handRank) -1 {
		if i == 0 {
			leftWins = 0
		}
		left := handRank[i]
		right := handRank[j]
		if v := WhichWins(left, right); v == 0 {
			leftWins++
		} else {
			handRank[i], handRank[j] = handRank[j], handRank[i]
		}
		if j == len(handRank)-1 {
			i = 0
			j = 1
		} else {
			i++
			j++
		}
	}
}