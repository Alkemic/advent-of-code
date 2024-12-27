package org.aoe.y2024.day_25

import java.io.File

fun main(args: Array<String>) {
    val input = File(args.toList().first()).readText()

    println("pt1: ${pt1(input)}")
}

fun pt1(input: String): Int {
    val (locksGroup, keysGroup) = input.split("\n\n").partition { it.startsWith("#####") }

    val parse = { str: String ->
        val lines = str.trim().lines()
        (0..4).map { i -> lines.count { line -> line[i] == '#' } }
    }

    val locks = locksGroup.map(parse)
    val keys = keysGroup.map(parse)
    return locks.sumOf { lock ->
        keys.count { key ->
            lock.zip(key) { l, k -> l + k }.all { it <= 7 }
        }
    }
}
