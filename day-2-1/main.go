package main

import (
    "bufio"
    "fmt"
    "github.com/dlclark/regexp2"
    "os"
    "strconv"
    "strings"
)

func SplitGameIDAndSubsets(s string) (int, []string) {
    lastDigitIntOrStringMatch, _ := regexp2.Compile("Game\\s(\\d+):\\s(.*)", 0)
    match, _ := lastDigitIntOrStringMatch.FindStringMatch(s)
    groups := match.Groups()
    gameID, _ := strconv.Atoi(groups[1].Captures[0].String())
    return gameID, strings.Split(groups[2].Captures[0].String(), "; ")
}


func CheckSubset(subset []string) bool {

    validSubsets := map[string]int {"red": 12, "green": 13, "blue": 14}

    for _, choices := range subset {
        choice := strings.Split(choices, ", ")

        for _, selection := range choice {

            amount, _ := strconv.Atoi(strings.Split(selection, " ")[0])
            color := strings.Split(selection, " ")[1]

            if validSubsets[color] < amount {
                return false
            }
        }
    }

    return true
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

        gameID, subsets := SplitGameIDAndSubsets(textLine)

        gameIsValid := CheckSubset(subsets)

        if gameIsValid {
            totalSum += gameID
        }

    }

    fmt.Println(totalSum)

    readFile.Close()
}
