package main

import (
	"os"
	"slices"
)

func main() {
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	println("pt2", pt2(string(input)))
}

func prepare(in string) []int {
	var out []int
	for i, ch := range in {
		n := int(ch - 48)
		if i%2 == 0 {
			out = append(out, slices.Repeat([]int{i / 2}, n)...)
		} else {
			out = append(out, slices.Repeat([]int{-1}, n)...)
		}
	}

	return out
}

func findFreeBlocks(disk []int, from, length int) (start int, ok bool) {
	for i, j := from, 0; i < len(disk) && i < len(disk); i++ {
		if disk[i] != -1 {
			continue
		}
		for j = i; j < len(disk) && disk[j] == -1; j++ {
		}

		if j-i >= length {
			return i, true
		}
		i = j
	}

	return -1, false
}

func pt2(in string) int {
	disk := prepare(in)

	lastMovedID, currentMovedID := disk[len(disk)-2]+1, disk[len(disk)-2]
	for j := len(disk) - 1; j >= 0; j-- {
		startJ := j + 1
		if disk[j] == -1 {
			continue
		}
		for ; j > 0 && disk[j] == disk[j-1]; j-- {
		}

		currentMovedID = disk[j]
		if lastMovedID < currentMovedID {
			continue
		}

		length := startJ - j
		// find empty enough space from begining
		startI, ok := findFreeBlocks(disk, 0, length)
		if !ok || startI >= j {
			continue
		}

		// swap
		for k := 0; k < length; k++ {
			disk[startI+k], disk[j+k] = disk[j+k], disk[startI+k]
		}
		lastMovedID = currentMovedID
	}

	return checksum(disk)
}

func checksum(in []int) int {
	sum := 0
	for i, n := range in {
		if n < 0 {
			continue
		}
		sum += n * i
	}

	return sum
}
