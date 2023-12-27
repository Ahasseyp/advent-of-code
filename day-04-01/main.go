package main

import (
    "fmt"
    "math"
    "regexp"
    "slices"
    "strconv"

    "advent-of-code/utils"
)

func extractNumbers(textLine string) []int {
    numberExpression := regexp.MustCompile("\\d+")

    var numbers []int

    for _, stringNumber := range numberExpression.FindAllString(textLine, -1) {
        number, _ := strconv.Atoi(stringNumber)
        numbers = append(numbers, number)
    } 

    return numbers
}

func parseTextLine(textLIne string) ([]int, []int){
    charExpression := regexp.MustCompile(`Card\s+\d+:\s+([\d\s]+)\|\s+([\d\s]+)`)
    matches := charExpression.FindStringSubmatch(textLIne)
    winningNumbers := extractNumbers(matches[1])
    selectedNumbers := extractNumbers(matches[2])
    return winningNumbers, selectedNumbers
}

func main()  {

    text := utils.ReadFile("./input.txt")

    totalPoints := 0

    for _, textLine := range text {
        winningNumbers, selectedNumbers := parseTextLine(textLine)
        slices.Sort(winningNumbers)
        slices.Sort(selectedNumbers)

        startingSelectedIndex := 0
        matchingNumberscount := 0

        for _, winningNumber := range winningNumbers {
            for i := startingSelectedIndex; i < len(selectedNumbers); i++ {
                selectedNumber := selectedNumbers[i]
                if winningNumber < selectedNumber {
                    startingSelectedIndex = i
                    break
                } else if winningNumber == selectedNumber {
                    matchingNumberscount += 1
                }
            }
        }

        totalPoints += int(math.Pow(2, float64(matchingNumberscount)-1))
    }

    fmt.Println(totalPoints)
}
