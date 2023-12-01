package main

import "fmt"

func main() {
	f := ParseFile("input.txt")
	first, err := FirstStar(f)
	
	if err != nil {
		fmt.Printf(`An error occured... somehow: %v`, err)
		return
	}
	fmt.Println(first)
}

func FirstStar(f []byte) (int, error) {
	res := 0
	numStore := []int{}

	for _, v := range f  {
		if num := int(v - 48); num > -1 && num < 10 {
			numStore = append(numStore, num)
		}

		if v == 10 {
			res += numStore[0] * 10 + numStore[len(numStore) - 1]
			numStore = []int{}
			continue
		}
	}

	return res, nil
}