package main

import (
	"fmt"
	"regexp"
	//"math"
    "slices"
	//"strconv"
    "strings"

	"advent-of-code/utils"
)

type rangeMap struct {
    source int
    destination int
    length int
}

func getMappedValue(ranges []rangeMap, sourceValue int) int {
    destinationValue := sourceValue

    for _, rangeMap := range ranges {
        if sourceValue >= rangeMap.source && sourceValue <= rangeMap.source + rangeMap.length - 1 {
            destinationValue = sourceValue + rangeMap.destination - rangeMap.source
        }
    }

    return destinationValue
}

func main()  {

    text := utils.ReadFile("./input.txt")

    seedsText, text := text[0], text[2:]

    numberExpression := regexp.MustCompile(`\d+`)

    seedsChar := numberExpression.FindAllString(seedsText, -1)

    seeds, _ := utils.StringSliceToIntegerSlice(seedsChar)

    input := strings.Join(text[:], "\n")

    var sourceToDestinationMap [][]rangeMap

    mapsExpression := regexp.MustCompile(`\w+-to-\w+\smap:\n([\d\s]+\d)`)

    matches := mapsExpression.FindAllStringSubmatch(input, -1)

    for _, match := range matches {

        var ranges []rangeMap
        lines := strings.Split(match[1], "\n")

        for _, line := range lines {

            rangeText := strings.Split(line, " ")

            data, _ := utils.StringSliceToIntegerSlice(rangeText)

            _range := rangeMap{destination: data[0], source: data[1], length: data[2]}

            ranges = append(ranges, _range)

        }

        sourceToDestinationMap = append(sourceToDestinationMap, ranges)
    }

    var results []int

    for _, seed := range seeds {
        for _, ranges := range sourceToDestinationMap {
            seed = getMappedValue(ranges, seed)
        }
        results = append(results, seed)
    }

    slices.Sort(results)

    fmt.Println(results[0])
}
