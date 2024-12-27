package aoe.y2024.d11

import java.io.File

fun main(args: Array<String>) {
    val input = File(args.toList().first()).readText().split(" ").map { it.toULong() }

    println("pt1: ${evolver(input, 25)}")
    println("pt2: ${evolver(input, 75)}")
}

val cache = mutableMapOf<Pair<ULong, Int>, ULong>()

fun evolver(stones: List<ULong>, blinks: Int): ULong {
    return stones.sumOf { count(it, blinks) }
}

fun count(stone: ULong, blinks: Int): ULong {
    return cache.getOrPut(stone to blinks) {
        when {
            blinks == 0 -> 1.toULong()
            stone == 0.toULong() -> count(1.toULong(), blinks - 1)
            stone.toString().length % 2 == 0 -> {
                count(stone.toString().drop(stone.toString().length / 2).toULong(), blinks - 1) +
                        count(stone.toString().dropLast(stone.toString().length / 2).toULong(), blinks - 1)
            }

            else -> count(stone * 2024.toULong(), blinks - 1)
        }
    }
}