package main

import (
	"bytes"
	"os"
)

func main() {
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	println("antinodes locations:", countAntinodes(bytes.Split(input, []byte{'\n'})))
}

type pos struct {
	x, y int
}

func (p pos) sub(x pos) pos {
	return pos{x: p.x - x.x, y: p.y - x.y}
}

func (p pos) add(x pos) pos {
	return pos{x: p.x + x.x, y: p.y + x.y}
}

func (p pos) withinBonds(width, height int) bool {
	return p.x < width && p.x >= 0 && p.y >= 0 && p.y < height
}

func countAntinodes(antenaMap [][]byte) int {
	// antennas by signal type
	antennas := map[byte][]pos{}
	for y, row := range antenaMap {
		for x, field := range row {
			if field != '.' {
				antennas[field] = append(antennas[field], pos{x, y})
			}
		}
	}

	mapWidth, mapHeight := len(antenaMap[0]), len(antenaMap)
	antinodes := map[pos]struct{}{}
	for _, pos := range antennas {
		for _, antinode := range calculateAntiNodes(pos, mapWidth, mapHeight) {
			// potential antinode must be within map and be empty spot
			if !antinode.withinBonds(mapWidth, mapHeight) {
				continue
			}

			antinodes[antinode] = struct{}{}
		}
	}

	return len(antinodes)
}

func calculateAntiNodes(positions []pos, width, height int) []pos {
	var antiNodes []pos
	for i, pos1 := range positions {
		for _, pos2 := range positions[i+1:] {
			deltaX, deltaY := pos1.x-pos2.x, pos1.y-pos2.y
			for j := width * -1; j < width; j++ { // bruteforcing this part...
				antiNodes = append(antiNodes,
					pos{x: pos1.x + deltaX*j, y: pos1.y + deltaY*j},
					pos{x: pos2.x - deltaX*j, y: pos2.y - deltaY*j},
				)
			}
		}
	}
	return antiNodes
}
