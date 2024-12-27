package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	theMap := bytes.Split(input, []byte{'\n'})
	println("trailhead count:", trailheads(theMap))
	println("trailhead2 count:", trailheads2(theMap))

}

var (
	neighboursDeltas = []pos{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
)

type twoDMap [][]byte

func (m twoDMap) Neighbours(p pos) (neighbours []pos) {
	for _, delta := range neighboursDeltas {
		n := pos{p.y - delta.y, p.x - delta.x}
		if m.withinBonds(n) {
			neighbours = append(neighbours, n)
		}
	}
	return neighbours
}

func (m twoDMap) withinBonds(p pos) bool {
	return p.x < len(m[0]) && p.x >= 0 && p.y >= 0 && p.y < len(m)
}

func (m twoDMap) drawpath(path []pos) {
	var drawMap [][]byte
	for i := 0; i < len(m); i++ {
		drawMap = append(drawMap, slices.Repeat([]byte{'.'}, len(m[0])))
	}
	for _, point := range path {
		drawMap[point.y][point.x] = m[point.y][point.x]
	}
	fmt.Println(string(bytes.Join(drawMap, []byte{'\n'})))
}

type pos struct{ y, x int }

func trailheads(theMap [][]byte) int {
	var starts []pos
	for y, row := range theMap {
		for x, ch := range row {
			if ch == '0' {
				starts = append(starts, pos{y, x})
			}
		}
	}

	sum := 0
	for _, start := range starts {
		ends := map[string]struct{}{}
		for _, path := range aStar(theMap, []pos{start}) {
			ends[fmt.Sprint(path[9])] = struct{}{}
		}
		sum += len(ends)
	}

	return sum
}

func trailheads2(theMap [][]byte) int {
	var starts []pos
	for y, row := range theMap {
		for x, ch := range row {
			if ch == '0' {
				starts = append(starts, pos{y, x})
			}
		}
	}
	return len(aStar(theMap, starts))
}

func aStar(theMap twoDMap, starts []pos) [][]pos {
	var paths [][]pos
	for _, start := range starts {
		if theMap[start.y][start.x] == '9' {
			paths = append(paths, []pos{start})
			continue
		}

		var validNeighbours []pos
		for _, n := range theMap.Neighbours(start) {
			if theMap[n.y][n.x] == theMap[start.y][start.x]+1 { // if at n is start+1
				validNeighbours = append(validNeighbours, n)
			}
		}

		for _, subPaths := range aStar(theMap, validNeighbours) {
			paths = append(paths, append([]pos{start}, subPaths...))
		}
	}

	return paths
}
