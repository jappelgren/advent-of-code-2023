package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Welcome to day 6")
	f := ParseFileToStringByNewLine("input.txt")

	first := FirstStar(f)
	fmt.Printf("First Star Result: %v%v", first, "\r\n")

	second := SecondStar(f)
	fmt.Printf("Second Star Result: %v%v", second, "\r\n")
}

func FirstStar(f []string) (res int) {
	times := []int{}
	records := []int{}

	for i, v := range f {
		noLabel := regexp.MustCompile("[a-zA-Z\\:]{1,}\\s{1,}").ReplaceAll([]byte(v), []byte(""))
		seperated := regexp.MustCompile("\\s{1,}").Split(string(noLabel), -1)
		for _, w := range seperated {
			var p *[]int
			if i == 0 {
				p = &times
			} else {
				p = &records
			}
			num, err := strconv.Atoi(w)
			if err == nil {
				*p = append(*p, num)
			}
		}
	}

	wins := 0
	for i, v := range times {
		for j := 1; j < v; j++ {
			if (v-j)*j > records[i] {
				wins++
			}
		}
		if res > 0 {
			res = res * wins
		} else {
			res = wins
		}
		wins = 0
	}
	return
}

func SecondStar(f []string) (res int) {
	var times int
	var records int

	for i, v := range f {
		noLabel := regexp.MustCompile("[a-zA-Z\\:]{1,}\\s{1,}").ReplaceAll([]byte(v), []byte(""))
		numStr := regexp.MustCompile("\\s{1,}").ReplaceAll(noLabel, []byte(""))
		num, err := strconv.Atoi(string(numStr))

		if i == 0 && err == nil {
			times = num
		} else if err == nil {
			records = num
		}

	}

	for i := 1; i < times; i++ {
		if (times-i)*i > records {
			res++
		}
	}

	return
}
