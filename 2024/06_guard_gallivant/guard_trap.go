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

func deepCopy(src [][]byte) [][]byte {
	dst := make([][]byte, len(src))
	for i := range src {
		dst[i] = make([]byte, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	labMapOrig := bytes.Split(input, []byte{'\n'})
	labMap := deepCopy(labMapOrig)

	cyclicCount := 0
	guardPosition := findGuard(labMap)
	// store locations in map[struct{x, y int}]byte, for easier counting, instead of going through labMap
	visited, _ := getPath(labMap, guardPosition)
	println("visited", visited)
	for position, _ := range visited {
		labMapCopy := deepCopy(labMapOrig)
		labMapCopy[position.y][position.x] = obstacle
		//printMap(labMapCopy)
		if _, cyclic := getPath(labMap, guardPosition); cyclic {
			cyclicCount++
		}
		//printMap(labMapCopy)
	}

	println("cyclic trap:", cyclicCount) // +1 for last move, guard didn't do
}

func nextPosition(labMap [][]byte, guardPosition pos) pos {
	current := labMap[guardPosition.y][guardPosition.x]
	switch current {
	case guardUp:
		return pos{x: guardPosition.x, y: guardPosition.y - 1}
	case guardRight:
		return pos{x: guardPosition.x + 1, y: guardPosition.y}
	case guardDown:
		return pos{x: guardPosition.x, y: guardPosition.y + 1}
	case guardLeft:
		return pos{x: guardPosition.x - 1, y: guardPosition.y}
	}
	return pos{}
}

func getPath(labMap [][]byte, guardPosition pos) (map[pos]byte, bool) {
	// position => direction at given position
	visited := map[pos]byte{}
	//printMap(labMap)

	for guardPosition, ok := guardMove(labMap, guardPosition); ok; guardPosition, ok = guardMove(labMap, guardPosition) {
		// check for cyclicity, if next position will result in guard being at tile that visited before
		nextGuardPosition := nextPosition(labMap, guardPosition)
		if visited[pos{y: nextGuardPosition.y, x: nextGuardPosition.x}] == labMap[guardPosition.y][guardPosition.x] {
			fmt.Println("lab map cyclic")
			printMap(labMap)
			return visited, true
		}

		//printMap(labMap)
		visited[pos{y: guardPosition.y, x: guardPosition.x}] = labMap[guardPosition.y][guardPosition.x]

	}
	fmt.Println("lab map")
	printMap(labMap)
	return visited, false
}

func withinBounds(x, y, width, height int) bool {
	width--
	height--
	//println(x, y, width, height)
	//println(y < 0, y > height, x < 0, x > width)
	return !(y < 0 || y > height || x < 0 || x > width)
}

func printMap(labMap [][]byte) {
	for _, row := range labMap {
		fmt.Println(string(row))
	}
}

// guardMove moves guard, and returns next position and if within lab
func guardMove(labMap [][]byte, curr pos) (pos, bool) {
	//defer func() {
	//	// lazy coding :-D
	//	if r := recover(); r != nil {
	//		//println("xvisited:", len(visited))
	//		fmt.Println(r)
	//	}
	//}()

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
		println(x, y+1, width, height)
		printMap(labMap)
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
