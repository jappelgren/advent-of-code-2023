package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	input := ParseFileToStringByNewLine("input.txt")
	fmt.Println("Welcome to day 5")

	first := FirstStar(input, InitSeedToSoil)
	fmt.Printf(`First Star result: %v%v`, first, "\r\n")

	second := SecondStar(input, InitSeedToSoilV2)
	fmt.Printf(`Second Star result: %v%v`, second, "\r\n")
}

type initFunc func(string, map[int]int)

func FirstStar(f []string, init initFunc) (res uint64) {
	res = ^uint64(0)

	seedToSoil := map[int]int{}
	soilToFertilizer := map[int]int{}
	fertilizerToWater := map[int]int{}
	waterToLight := map[int]int{}
	lightToTemp := map[int]int{}
	tempToHumidity := map[int]int{}
	humidityToLocation := map[int]int{}

	curMap := seedToSoil
	nextMap := soilToFertilizer

	for i, v := range f {
		if i == 0 {
			init(v, seedToSoil)
			continue
		}

		if v == "" {
			continue
		}

		switch v {
		case "seed-to-soil map:":
			break
		case "soil-to-fertilizer map:":
			curMap = soilToFertilizer
			nextMap = fertilizerToWater
			break
		case "fertilizer-to-water map:":
			curMap = fertilizerToWater
			nextMap = waterToLight
			break
		case "water-to-light map:":
			curMap = waterToLight
			nextMap = lightToTemp
			break
		case "light-to-temperature map:":
			curMap = lightToTemp
			nextMap = tempToHumidity
			break
		case "temperature-to-humidity map:":
			curMap = tempToHumidity
			nextMap = humidityToLocation
			break
		case "humidity-to-location map:":
			curMap = humidityToLocation
			nextMap = nil
			break
		default:
			mapVals := regexp.MustCompile("\\s").Split(v, -1)
			var destFloor int
			var sourceFloor int
			var sourceCeil int
			var valRange int

			for j, w := range mapVals {
				if x, err := strconv.Atoi(w); err == nil {
					switch j {
					case 0:
						destFloor = x
					case 1:
						sourceFloor = x
					case 2:
						valRange = x
						sourceCeil = sourceFloor + valRange - 1
					}
				}
			}
			for k := range curMap {
				if k >= sourceFloor && k <= sourceCeil {
					toMapVal := k - sourceFloor + destFloor
					curMap[k] = toMapVal
				}
				if nextMap != nil {
					nextMap[curMap[k]] = curMap[k]
				}
			}
		}
	}

	for _, v := range seedToSoil {
		seedToLocation := humidityToLocation[tempToHumidity[lightToTemp[waterToLight[fertilizerToWater[soilToFertilizer[v]]]]]]
		if uint64(seedToLocation) < res {
			res = uint64(seedToLocation)
		}
	}
	return
}

func SecondStar(f []string, init initFunc) (res uint64) {
	res = FirstStar(f, init)
	return
}

func InitSeedToSoil(v string, seedToSoil map[int]int) {
	seedIdsStr := regexp.MustCompile("seeds:\\s").ReplaceAll([]byte(v), []byte(""))
	seedIds := regexp.MustCompile("\\s").Split(string(seedIdsStr), -1)
	for _, w := range seedIds {
		if num, err := strconv.Atoi(w); err == nil {
			seedToSoil[num] = num
		}
	}
}

func InitSeedToSoilV2(v string, seedToSoil map[int]int) {
	seedIdsStr := regexp.MustCompile("seeds:\\s").ReplaceAll([]byte(v), []byte(""))
	seedIds := regexp.MustCompile("\\s").Split(string(seedIdsStr), -1)
	for i := 0; i < len(seedIds); i += 2 {
		floor := 0
		ceil := 0
		if num, err := strconv.Atoi(seedIds[i]); err == nil {
			floor = num
		}
		if num, err := strconv.Atoi(seedIds[i+1]); err == nil {
			ceil = floor + num - 1
		}
		for j := floor; j <= ceil; j++ {
			seedToSoil[j] = j
		}
	}

}
