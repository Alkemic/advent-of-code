package org.aoe.day_08

import java.io.File
import kotlin.math.max

data class Pos(val x: Int, val y: Int) {
    operator fun plus(other: Pos) = Pos(x + other.x, y + other.y)
    operator fun minus(other: Pos) = Pos(x - other.x, y - other.y)
    operator fun times(other: Int) = Pos(x * other, y * other)

    fun withinMap(height: Int, width: Int): Boolean {
        return x in 0..<width && y in 0..<height
    }
}

fun main(args: Array<String>) {
    val input = File(args.toList().first())
    val map = input.readLines().map { it.toList() }

    println("pt1: ${findAntiNodes(map, false)}")
    println("pt2: ${findAntiNodes(map, true)}")
}

fun findAntiNodes(map: List<List<Char>>, withHarmonics: Boolean): Int {
    val antennas = map.map { it.toList() }
        .flatMapIndexed { y, row ->
            row.mapIndexedNotNull { x, c ->
                if (c != '.') c to Pos(x, y) else null
            }
        }
        .groupingBy { it.first }
        .fold(emptyList<Pos>()) { acc, (_, value) -> acc + value }

    val mapHeight = map.first().size
    val mapWidth =  map.first().size
    val antiNodesLoc = mutableSetOf<Pos>()
    antennas.forEach({ (ant, positions) ->
        positions.forEachIndexed{ i, pos1 ->
            positions.subList(i+1, positions.size).forEach { pos2 ->
                val delta = pos1-pos2
                if (withHarmonics) {
                    // bruteforcing...
                    (-max(mapHeight, mapWidth)..max(mapHeight, mapWidth)).forEach { i ->
                        antiNodesLoc += pos1+delta*i
                        antiNodesLoc += pos2-delta
                    }
                } else {
                    antiNodesLoc += pos1+delta
                    antiNodesLoc += pos2-delta
                }
            }
        }
    })

    val validAntiNode = antiNodesLoc.mapNotNull { if (it.withinMap(mapHeight, mapWidth)) it else null }

    return validAntiNode.size
}