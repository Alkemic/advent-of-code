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

	antinodes := map[pos]struct{}{}
	for _, pos := range antennas {
		for _, antinode := range calculateAntiNodes(pos) {
			// potential antinode must be within map and be empty spot
			if !(antinode.withinBonds(len(antenaMap[0]), len(antenaMap))) {
				continue
			}

			antinodes[antinode] = struct{}{}
		}
	}

	return len(antinodes)
}

func calculateAntiNodes(positions []pos) []pos {
	var antiNodes []pos
	for i, pos1 := range positions {
		for _, pos2 := range positions[i+1:] {
			delta := pos1.sub(pos2)
			//deltaX, deltaY := pos1.x-pos2.x, pos1.y-pos2.y
			//add to pos1 and subtract from pos2
			antiNodes = append(antiNodes, pos1.add(delta), pos2.sub(delta))
			//antiNodes = append(antiNodes, pos{x: pos1.x + deltaX, y: pos1.y + deltaY}, pos{x: pos2.x - deltaX, y: pos2.y - deltaY})
		}
	}
	return antiNodes
}
