package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to day 9")
	f := ParseFileToBytes("input.txt")

	first := FirstStar(f)
	fmt.Printf("First Star Result: %v%v", first, "\r\n")

	second := SecondStar(f)
	fmt.Printf("Second Star Result: %v%v", second, "\r\n")
}

func FirstStar(f []byte) (res int) {
	startIndex, rowLen, input := ParseInput(f)
	stepCount := 0
	var nextStep int
	var lastStep int
	var nextDir string
	opMap := map[string]string{
		"UP":    "DOWN",
		"DOWN":  "UP",
		"LEFT":  "RIGHT",
		"RIGHT": "LEFT",
	}

	if stepCount == 0 {
		for k := range opMap {
			s, d := CheckSurroundings(input, startIndex, rowLen, k, opMap)
			if s >= 0 && input[s] != '.' {
				lastStep = startIndex
				nextStep = s
				nextDir = d
				break
			}
		}
		stepCount = 1
	}

	for true {
		s, d := CheckSurroundings(input, nextStep, rowLen, nextDir, opMap)
		if s >= 0 && s != lastStep {
			lastStep = nextStep
			nextStep = s
			nextDir = d
		}
		stepCount++
		if input[nextStep] == 'S' {
			res = stepCount / 2
			break
		}
	}
	return
}

func SecondStar(f []byte) (res int) {
	return
}

func ParseInput(f []byte) (startIndex int, rowLen int, input []byte) {
	startIndex = -1
	rowLen = -1
	for i, v := range f {
		if v == '\r' || v == '\n' {
			if rowLen < 0 {
				rowLen = i
			}
		} else {
			input = append(input, byte(v))
		}
	}
	for i, v := range input {
		if v == 'S' {
			startIndex = i
			break
		}
	}
	return
}

func CheckSurroundings(pipeMap []byte, position int, rowLen int, direction string, opMap map[string]string) (nextPos int, nextDir string) {
	dirMap := GetDirMap()
	nextPos = -1
	switch direction {
	case "UP":
		if position >= rowLen {
			acceptableOptions := dirMap[pipeMap[position]]["UP"]
			for _, v := range acceptableOptions {
				if v == pipeMap[position-rowLen] {
					nextPos = position - rowLen
					break
				}
			}
		}
	case "DOWN":
		if len(pipeMap)-rowLen > position {
			acceptableOptions := dirMap[pipeMap[position]]["DOWN"]
			for _, v := range acceptableOptions {
				if v == pipeMap[position+rowLen] {
					nextPos = position + rowLen
					break
				}
			}
		}
	case "LEFT":
		if position%rowLen > 0 {
			acceptableOptions := dirMap[pipeMap[position]]["LEFT"]
			for _, v := range acceptableOptions {
				if v == pipeMap[position-1] {
					nextPos = position - 1
					break
				}
			}
		}
	case "RIGHT":
		if position%rowLen+1 < rowLen {
			acceptableOptions := dirMap[pipeMap[position]]["RIGHT"]
			for _, v := range acceptableOptions {
				if v == pipeMap[position+1] {
					nextPos = position + 1
					break
				}
			}
		}
	}
	if nextPos >= 0 {
		for k := range dirMap[pipeMap[nextPos]] {
			if k != opMap[direction] {
				nextDir = k
			}
		}
	}
	return
}

func GetDirMap() map[byte]map[string][]byte {
	return map[byte]map[string][]byte{
		'|': {
			"UP":   []byte{'7', 'F', '|', 'S'},
			"DOWN": []byte{'L', 'J', '|', 'S'},
		},
		'-': {
			"LEFT":  []byte{'L', 'F', '-', 'S'},
			"RIGHT": []byte{'J', '7', '-', 'S'},
		},
		'L': {
			"UP":    []byte{'|', '7', 'F', 'S'},
			"RIGHT": []byte{'7', '-', 'J', 'S'},
		},
		'J': {
			"UP":   []byte{'|', 'F', '7', 'S'},
			"LEFT": []byte{'-', 'F', 'L', 'S'},
		},
		'7': {
			"DOWN": []byte{'|', 'L', 'J', 'S'},
			"LEFT": []byte{'-', 'F', 'L', 'S'},
		},
		'F': {
			"DOWN":  []byte{'|', 'J', 'L', 'S'},
			"RIGHT": []byte{'-', 'J', '7', 'S'},
		},
		'S': {
			"UP":    []byte{'|', '7', 'F'},
			"DOWN":  []byte{'|', 'J', 'L'},
			"LEFT":  []byte{'-', 'L', 'F'},
			"RIGHT": []byte{'-', 'J', '7'},
		},
	}
}
