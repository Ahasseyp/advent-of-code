package main

import (
    "fmt"
    "strings"
    "regexp"
    "strconv"

    "advent-of-code/utils"
)

func padSlice(slice []string, char string) []string {
    horizontalPaddingSlice := []string{strings.Repeat(char, len(slice[0]))}
    slice = append(horizontalPaddingSlice, slice...)
    slice = append(slice, horizontalPaddingSlice...)

    for i, line := range slice {
        slice[i] = "." + line + "."
    }

    return slice
}

func checkIsAdjecent(slice []string, x int, y int) bool {
    charExpression := regexp.MustCompile("[^\\d^\\.]")

    for diffX := -1; diffX < 2; diffX++ {
        for diffY := -1; diffY < 2; diffY++ {
            if charExpression.MatchString(string(slice[x+diffX][y+diffY])) {
                return true
            }
        }
    }

    return false
}

func main()  {

    text := utils.ReadFile("./input.txt")

    text = padSlice(text, ".")

    matches := make(map[string][][2]int)

    for lineNumber, textLine := range text {
        numberMatchExpression := regexp.MustCompile("\\d+")

        numberMatch := numberMatchExpression.FindAllString(textLine, -1)
        numberMatchIndex := numberMatchExpression.FindAllStringSubmatchIndex(textLine, -1)

        for i := range numberMatch {

            index := fmt.Sprintf("%s,%d,%d", numberMatch[i], lineNumber, numberMatchIndex[i][0])

            for y := numberMatchIndex[i][0]; y < numberMatchIndex[i][1]; y++ {
                coordintes := [2]int{lineNumber, y}
                matches[index] = append(matches[index], coordintes)
            }
        }
    }

    totalSum := 0

    for match, coordinates := range matches {

        number := strings.Split(match, ",")

        matchIsAdjecent := false

        for _, coordinate := range coordinates {
            if checkIsAdjecent(text, coordinate[0], coordinate[1]) {
                matchIsAdjecent = true
            }
        }

        if matchIsAdjecent {
            intMatch, _ := strconv.Atoi(number[0])
            totalSum += intMatch
        }
    }

    fmt.Println(totalSum)
}
