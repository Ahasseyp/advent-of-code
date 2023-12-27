package main

import (
    "bufio"
    "fmt"
    "os"
    "github.com/dlclark/regexp2"
    "strconv"
)

func getFirstDigit(s string) string {
    firstDigitMatch, _ := regexp2.Compile("\\d", 0)
    match, _ := firstDigitMatch.FindStringMatch(s)
    return match.String()
}

func getLastDigit(s string) string {
    lastDigitMatch, _ := regexp2.Compile("(\\d{1})(?!.*\\d)", 0)
    match, _ := lastDigitMatch.FindStringMatch(s)
    return match.String()
}

func main()  {
    readFile, err := os.Open("./input.txt")

    if err != nil {
        fmt.Println(err)
    }

    fileScanner := bufio.NewScanner(readFile)

    fileScanner.Split(bufio.ScanLines)

    var totalSum int

    totalSum = 0

    for fileScanner.Scan() {
        textLine := fileScanner.Text()

        firstDigit := getFirstDigit(textLine)
        lastDigit := getLastDigit((textLine))

        sum, _ := strconv.Atoi(firstDigit+lastDigit)

        totalSum += sum
    }

    fmt.Println(totalSum)

    readFile.Close()
}
