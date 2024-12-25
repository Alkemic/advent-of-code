package org.aoe.day_07

import java.io.File
import kotlin.math.pow

fun main(args: Array<String>) {
    val input = File(args.toList().first()).readText()

    val add = { a: ULong, b: ULong -> a + b }
    val mul = { a: ULong, b: ULong -> a * b }
    val concat = { a: ULong, b: ULong -> (a.toString() + b.toString()).toULong() }

    println("pt1: ${calibration(input, arrayOf(add, mul))}")
    println("pt2: ${calibration(input, arrayOf(add, mul, concat))}")
}

fun calibration(input: String, ops: Array<(ULong, ULong) -> ULong>): ULong {
    val calibrations = input.lines()
        .map { line ->
            val parts = line.split(": ")
            parts.first().toULong() to parts.last().split(" ").map { it.toULong() }
        }

    // todo: use fold
    val sum = calibrations.mapNotNull { (x, nums) ->
        val base = ops.size
        val results = mutableListOf(nums.first())
        var a = 0
        var b = 0
        for (i in 1..<nums.size) {
            for (s in a..b) {
                results.addAll(ops.map { f -> f(results[s], nums[i]) })
            }
            a += base.toDouble().pow(i - 1).toInt()
            b += base.toDouble().pow(i).toInt()
        }

        if (x in results.slice((results.size / base)..<results.size)) x else null
    }

    return sum.sum()
}
