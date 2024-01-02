package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"advent-of-code/utils"
)

type Galaxy struct {
    index int
    x, y int
}

func flipInput(text []string) []string {
	result := []string{}

	for y := 0; y < len(text[0]); y++ {
		var column string
		for x := 0; x < len(text); x++ {
			column += string(text[x][y])
		}
		result = append(result, column)
	}

	return result
}

func getRowsWithNoGalaxies(text []string) []int {
	result := []int{}

	for index, textLine := range text {
		if !strings.Contains(textLine, "#") {
			result = append(result, index)
		}
	}

	return result
}

func getColumnsWithNoGalaxies(text []string) []int {
	flippedInput := flipInput(text)

	return getRowsWithNoGalaxies(flippedInput)
}

func addNewRows(text []string, rowsWithNoGalaxies []int) []string {
	result := []string{}

	for index, textLine := range text {
		result = append(result, textLine)
		if slices.Contains(rowsWithNoGalaxies, index) {
			result = append(result, strings.Repeat(".", len(textLine)))
		}
	}

	return result
}

func addNewColumns(text []string, columnsWithNoGalaxies []int) []string {
    text = flipInput(text)
    text = addNewRows(text, columnsWithNoGalaxies)
    text = flipInput(text)

    return text
}

func expandInput(text []string, rowsWithNoGalaxies, columnsWithNoGalaxies []int) []string {
	text = addNewRows(text, rowsWithNoGalaxies)
    text = addNewColumns(text, columnsWithNoGalaxies)

	return text
}

func getGalaxies(text []string) []Galaxy {
    galaxies := []Galaxy{}

    for y, row := range text {
        for x, char := range row {
            if string(char) == "#" {
                galaxies = append(galaxies, Galaxy{index: len(galaxies)+1, x: x, y: y})
            }
        }
    }

    return galaxies
}

func distanceBetweenGalaxies(galaxyOne, galaxyTwo Galaxy) (distance int) {
    horizontalDistance := int(math.Abs(float64(galaxyOne.x) - float64(galaxyTwo.x)))
    verticalDistance := int(math.Abs(float64(galaxyOne.y) - float64(galaxyTwo.y)))
    return horizontalDistance + verticalDistance
}

func main() {
	text := utils.ReadFile("./input.txt")

    rowsWithNoGalaxies := getRowsWithNoGalaxies(text)
    columnsWithNoGalaxies := getColumnsWithNoGalaxies(text)

    text = expandInput(text, rowsWithNoGalaxies, columnsWithNoGalaxies)

    galaxies := getGalaxies(text)

    var result int

    for index, galaxyOne := range galaxies {
        for _, galaxyTwo := range galaxies[index+1:] {
            distance := distanceBetweenGalaxies(galaxyOne, galaxyTwo)
            result += distance
        }
    }

    fmt.Println(result)
}
