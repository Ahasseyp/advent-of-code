package main

import (
    "bufio"
    "fmt"
    "github.com/dlclark/regexp2"
    "os"
    "strconv"
    "strings"
)

func SplitGameIDAndSubsets(s string) []string {
    lastDigitIntOrStringMatch, _ := regexp2.Compile("Game\\s(\\d+):\\s(.*)", 0)
    match, _ := lastDigitIntOrStringMatch.FindStringMatch(s)
    groups := match.Groups()
    return strings.Split(groups[2].Captures[0].String(), "; ")
}

func MinimumSubset(subset []string) map[string]int {

    minimumSubsets := map[string]int {"red": 0, "green": 0, "blue": 0}

    for _, choices := range subset {
        choice := strings.Split(choices, ", ")

        for _, selection := range choice {

            amount, _ := strconv.Atoi(strings.Split(selection, " ")[0])
            color := strings.Split(selection, " ")[1]

            if amount > minimumSubsets[color] {
                minimumSubsets[color] = amount
            }
            
        }
    }

    return minimumSubsets

}

func SubsetPower(subset map[string]int) int {

    result := 1

    for _, amount := range subset {
        result *= amount
    }

    return result
}

func main()  {
    readFile, err := os.Open("./input.txt")

    if err != nil {
    }

    fileScanner := bufio.NewScanner(readFile)

    fileScanner.Split(bufio.ScanLines)

    var totalSum int

    totalSum = 0

    for fileScanner.Scan() {
        textLine := fileScanner.Text()

        subsets := SplitGameIDAndSubsets(textLine)

        minimumSubset := MinimumSubset(subsets)

        totalSum += SubsetPower(minimumSubset)

    }

    fmt.Println(totalSum)

    readFile.Close()
}
