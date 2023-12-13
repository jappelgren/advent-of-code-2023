package main

import (
	"fmt"
	"regexp"

	"github.com/fxtlabs/primes"
)

type Node struct {
	l string
	r string
}

func main() {
	fmt.Println("Welcome to day 8")
	f := ParseFileToStringByNewLine("input.txt")

	first := FirstStar(f)
	fmt.Printf("First Star Result: %v%v", first, "\r\n")

	second := SecondStar(f)
	fmt.Printf("Second Star Result: %v%v", second, "\r\n")
}

func FirstStar(f []string) (res int) {
	mapMap, instructions := GetInstructionsAndMap(f)
	i := 0
	res = FindPath(mapMap, instructions, i)
	return
}

func SecondStar(f []string) (res int) {
	mapMap, instructions := GetInstructionsAndMap(f)
	aPaths := []string{}

	for k := range mapMap {
		if ll := []byte(k)[2]; ll == 'A' {
			aPaths = append(aPaths, k)
		}
	}
	
	shortestPaths := []int{}
	
	for _, v := range aPaths {
		shortestPaths = append(shortestPaths, FindPathV2(mapMap, instructions, v))
	} 
	
	lcmMap := map[int]bool{}
	
	for _, v := range shortestPaths  {
		primes := primes.Sieve(v) 
		for _, w := range primes {
			if r := v % w; r < 1 {
				lcmMap[w] = true
			}
		}
	}

	for k := range lcmMap {
		if res == 0 {
			res = k
		} else {
			res = res * k
		}
	}
	return
}

func GetInstructionsAndMap(f []string) (mapMap map[string]Node, instructions []string) {
	mapMap = map[string]Node{}
	instructions = regexp.MustCompile("").Split(f[0], -1)
	for i, v := range f {
		if i < 2 {
			continue
		}
		s := regexp.MustCompile("[\\s=\\(\\)\\,]{1,}").Split(v, -1)
		mapMap[string(s[0])] = Node{string(s[1]), string(s[2])}
	}

	for i, v := range f {
		if i < 2 {
			continue
		}
		s := regexp.MustCompile("[\\s=\\(\\)\\,]{1,}").Split(v, -1)
		mapMap[string(s[0])] = Node{string(s[1]), string(s[2])}
	}
	return
}

func FindPath(mapMap map[string]Node, instructions []string, i int) (res int) {
	nextNode := "AAA"
	for nextNode != "ZZZ" {
		res++
		dir := instructions[i]
		if dir == "L" {
			nextNode = mapMap[nextNode].l
		} else {
			nextNode = mapMap[nextNode].r
		}
		if i == len(instructions)-1 {
			i = 0
			continue
		}
		i++
	}
	return
}

func FindPathV2(mapMap map[string]Node, instructions []string, nextNode string) (res int) {
	i := 0
	for []byte(nextNode)[2] != 'Z' {
		res++
		dir := instructions[i]
		if dir == "L" {
			nextNode = mapMap[nextNode].l
		} else {
			nextNode = mapMap[nextNode].r
		}
		if i == len(instructions)-1 {
			i = 0
			continue
		}
		i++
	}
	return
}
