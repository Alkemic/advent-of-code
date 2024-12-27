package main

import (
	"fmt"
	math2 "math"
	"os"
	"strconv"
	"strings"

	"github.com/Alkemic/aoc/math"
)

var cache = make(map[string]int) // key: nums-blinks

func main() {
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	var nums []int
	for _, str := range strings.Split(string(input), " ") {
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}

	println("stones count 25:", evolver(nums, 25))
	println("stones count 75:", evolver(nums, 75))
}

func evolver(stones []int, blinks int) int {
	sum := 0
	for _, num := range stones {
		sum += count(num, blinks)
	}
	return sum
}

func count(stone, blinks int) int {
	key := fmt.Sprintf("%d-%d", stone, blinks)
	if sum, ok := cache[key]; ok {
		return sum
	}

	var sum int
	if blinks == 0 {
		cache[key] = 1
		return 1

	}
	if stone == 0 {
		sum = count(1, blinks-1)
		cache[key] = sum
		return sum
	}

	if len(fmt.Sprintf("%d", stone))%2 == 0 {
		half := math.Pow10((int(math2.Log10(float64(stone))) + 1) / 2)
		sum = count(stone/half, blinks-1) + count(stone%half, blinks-1)
		cache[key] = sum
		return sum
	}

	sum = count(stone*2024, blinks-1)
	cache[key] = sum
	return sum
}
