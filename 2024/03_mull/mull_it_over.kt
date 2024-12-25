package org.aoe.y2024.day_03

import java.io.File

fun main(args: Array<String>) {
    val input = File(args.toList().first()).readText().replace("\n", "")

    println("pt1: ${pt1(input)}")
    println("pt2: ${pt2(input)}")
}

const val mulPattern = """mul\(([0-9]{1,3}),([0-9]{1,3})\)"""
const val cutOutPattern = """(don't\(\).*)"""

val pt1 = { input: String ->
    mulPattern.toRegex().findAll(input).toList()
        .fold(0.toULong()) { t, r -> t + r.groupValues[1].toULong() * r.groupValues[2].toULong() }
}

val pt2 = { input: String ->
    pt1(
        input.replace("don't()", "\ndon't()").replace("do()", "\ndo()").replace(cutOutPattern.toRegex(), "")
            .replace("\n", "")
    )
}
