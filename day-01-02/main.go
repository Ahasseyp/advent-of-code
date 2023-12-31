package main

import (
    "bufio"
    "fmt"
    "os"
    "github.com/dlclark/regexp2"
    "strconv"
)

func getFirstDigitIntOrString(s string) string {
    firstDigitIntOrStringMatch, _ := regexp2.Compile("(\\d{1}|one|two|three|four|five|six|seven|eight|nine)", 0)
    match, _ := firstDigitIntOrStringMatch.FindStringMatch(s)
    if match != nil {
        return match.String()
    } else {
        return ""
    }
}

func getLastDigitIntOrString(s string) string {
    lastDigitIntOrStringMatch, _ := regexp2.Compile("(\\d{1}|one|two|three|four|five|six|seven|eight|nine)(?!.*(\\d|one|two|three|four|five|six|seven|eight|nine))", 0)
    match, _ := lastDigitIntOrStringMatch.FindStringMatch(s)
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

    mapping := map[string]string {
        "one": "1",
        "two": "2",
        "three": "3",
        "four": "4",
        "five": "5",
        "six": "6",
        "seven": "7",
        "eight": "8",
        "nine": "9",
    }

    for fileScanner.Scan() {
        textLine := fileScanner.Text()

        firstDigitIntOrString := getFirstDigitIntOrString(textLine)

        var lastDigitIntOrString string

        for i := 0; i < len(textLine); i++ {
            newTextLine := textLine[i:]
            match := getFirstDigitIntOrString((newTextLine))
            if match != "" {
                lastDigitIntOrString = match
            } else {
                break
            }
        }

        //fmt.Println(textLine, firstDigitIntOrString, lastDigitIntOrString)

        _, err := strconv.Atoi(firstDigitIntOrString)

        var firstDigitChar string

        if err != nil {
            firstDigitChar = mapping[firstDigitIntOrString]
        } else {
            firstDigitChar = firstDigitIntOrString
        }

        _, err1 := strconv.Atoi(lastDigitIntOrString)

        var lastDigitChar string

        if err1 != nil {
            lastDigitChar = mapping[lastDigitIntOrString]
        } else {
            lastDigitChar = lastDigitIntOrString
        }

        sum, err := strconv.Atoi(firstDigitChar+lastDigitChar)

        //fmt.Println(sum)

        totalSum += sum
    }

    fmt.Println(totalSum)

    readFile.Close()
}
