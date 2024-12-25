package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

const (
	guardUp    = '^'
	guardDown  = 'v'
	guardLeft  = '<'
	guardRight = '>'

	obstacle = '#'
	clear    = '.'
	visited  = 'X'
)

type pos struct{ y, x int }

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	labMap := bytes.Split(input, []byte{'\n'})

	guardPosition := findGuard(labMap)
	// store locations in map[struct{x, y int}]struct{}, for easier counting, instead of going through labMap
	visited := getPath(labMap, guardPosition)

	println("visited:", len(visited)+1) // +1 for last move, guard didn't do
}

func getPath(labMap [][]byte, guardPosition pos) map[pos]struct{} {
	visited := map[pos]struct{}{}
	//printMap(labMap)

	defer func() {
		// lazy coding :-D
		if r := recover(); r != nil {
			println("xvisited:", len(visited))
		}
	}()

	for guardPosition, ok := guardMove(labMap, guardPosition); ok; guardPosition, ok = guardMove(labMap, guardPosition) {
		//printMap(labMap)
		visited[pos{y: guardPosition.y, x: guardPosition.x}] = struct{}{}
	}

	return visited
}

func withinBounds(x, y, width, height int) bool {
	return !(y < 0 || y > height || x < 0 || x >= width)
}

func printMap(labMap [][]byte) {
	for _, row := range labMap {
		fmt.Println(string(row))
	}
}

// guardMove moves guard, and returns next position and if within lab
func guardMove(labMap [][]byte, curr pos) (pos, bool) {
	height := len(labMap)   // y coord
	width := len(labMap[0]) // x coord

	y := curr.y
	x := curr.x
	next := pos{y: y, x: x}

	current := labMap[curr.y][curr.x]
	switch current {
	case guardUp:
		if withinBounds(x, y-1, width, height) && labMap[y-1][x] == obstacle {
			next.x++
			labMap[next.y][next.x] = guardRight
		} else if withinBounds(x, y-1, width, height) {
			next.y--
			labMap[next.y][next.x] = current
		} else {
			return pos{-1, -1}, false
		}
	case guardRight:
		if withinBounds(x+1, y, width, height) && labMap[y][x+1] == obstacle {
			next.y++
			labMap[next.y][next.x] = guardDown
		} else if withinBounds(x+1, y, width, height) {
			next.x++
			labMap[next.y][next.x] = current
		} else {
			return pos{-1, -1}, false
		}
	case guardDown:
		if withinBounds(x, y+1, width, height) && labMap[y+1][x] == obstacle {
			next.x--
			labMap[next.y][next.x] = guardLeft
		} else if withinBounds(x, y+1, width, height) {
			next.y++
			labMap[next.y][next.x] = current
		} else {
			return pos{-1, -1}, false
		}
	case guardLeft:
		if withinBounds(x-1, y, width, height) && labMap[y][x-1] == obstacle {
			next.y--
			labMap[next.y][next.x] = guardUp
		} else if withinBounds(x+1, y, width, height) {
			next.x--
			labMap[next.y][next.x] = current
		} else {
			return pos{-1, -1}, false
		}
	}

	labMap[curr.y][curr.x] = visited

	return next, true
}

func findGuard(labMap [][]byte) pos {
	for y, labRow := range labMap {
		for x, tile := range labRow {
			if tile == obstacle || tile == clear {
				continue
			}

			return pos{y: y, x: x}
		}
	}

	panic("this shouldn't happens")
}
