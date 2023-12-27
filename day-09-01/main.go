package main

import (
	"fmt"
	"strings"

	"advent-of-code/utils"
)

func sliceContainsOnlyZeros(history []int) bool {
    for _, v := range history {
        if v != 0 {
            return false
        }
    }
    return true
}

func lastElement(slice []int) int {
    return slice[len(slice)-1]
}

func differenceForValues(history []int) []int {
    result := make([]int, 0)

    for i := 0; i < len(history)-1; i++ {
        difference := history[i+1] - history[i]
        result = append(result, difference)
    }

    return result
}

func nextElement(slice []int) int {
    if sliceContainsOnlyZeros(slice) {
        return 0
    } else {
        return lastElement(slice) + nextElement(differenceForValues(slice))
    }
}

func main() {
	text := utils.ReadFile("./input.txt")

    result := 0

	for _, textLine := range text {
        numbers := strings.Split(textLine, " ")
        history, _ := utils.StringSliceToIntegerSlice(numbers)
        result += nextElement(history)
	}

    fmt.Println(result)
}
