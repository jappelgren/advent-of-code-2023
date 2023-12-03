package main

import (
	"fmt"
	"strconv"
)

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

func main() {
	input := ParseFileToBytes("example.txt")
	fmt.Println("Welcome to day 3")

	first := FirstStar(input)
	fmt.Printf(`First Star result: %v%v`, first, "\r\n")

	second := SecondStar(input)
	fmt.Printf(`Second Star result: %v%v`, second, "\r\n")

}

func FirstStar(f []byte) (res int) {
	var rowLen int
	var noLBs []byte

	for i, v := range f {
		if v == 10 || v == 13 {
			if rowLen == 0 {
				rowLen = i
			}
			continue
		}
		noLBs = append(noLBs, v)
	}

	// adjacent := []int{}
	lastByte := byte(46)
	var numStr string
	indexMap := map[int][]int{}
	var tempISlice []int
	for i, v := range f {
		if num := v - 48; num >= 0 && num <= 9 {
			if len(numStr) == 0 {
				numStr = string(rune(num))
			} else if lastByte -48 >= 0 && lastByte - 48 <= 9 {
				numStr += string(rune(num)) 
			}
			tempISlice = append(tempISlice, i)
		} else {
			numKey, err := strconv.Atoi(numStr); if err == nil {
				indexMap[numKey] = tempISlice
				tempISlice = []int{}
				numStr = ""
			}
		}
		lastByte = v
	}
	fmt.Println(indexMap)

	return
}

func SecondStar(f []byte) (res int) {
	res = 0
	return
}

func FindAdjacentIndicies(i int, rowLen int, arLen int) (adjacentIndicies []int) {
	left := dir(i, rowLen, LEFT)
	right := dir(i, rowLen, RIGHT)
	below := dir(i, rowLen, DOWN)
	above := dir(i, rowLen, UP)
	aboveleft := dir(left, rowLen, UP)
	aboveright := dir(right, rowLen, UP)
	belowleft := dir(left, rowLen, DOWN)
	belowright := dir(right, rowLen, DOWN)

	// first row check
	if i < rowLen {
		above = -1
		aboveleft = -1
		aboveright = -1
	}
	// last row check
	if arLen-i <= rowLen {
		below = -1
		belowleft = -1
		belowright = -1
	}
	// first col check
	if i%rowLen == 0 {
		left = -1
		aboveleft = -1
		belowleft = -1
	}
	// last col check
	if i%rowLen+1 == rowLen {
		right = -1
		aboveright = -1
		belowright = -1
	}

	allDirs := []int{above, aboveright, right, belowright, below, belowleft, left, aboveleft}
	for _, v := range allDirs {
		if v >= 0 {
			adjacentIndicies = append(adjacentIndicies, v)
		}
	}

	return
}

func dir(i int, rowLen int, dir int) (dirI int) {
	switch dir {
	case UP:
		dirI = i - rowLen
	case RIGHT:
		dirI = i + 1
	case DOWN:
		dirI = i + rowLen
	case LEFT:
		dirI = i - 1
	}
	return
}
