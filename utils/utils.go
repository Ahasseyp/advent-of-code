package utils

import (
	"bufio"
	"os"
	"strconv"
    "github.com/dlclark/regexp2"
)

// Helper function to read a file
// Input is a file path
// Output is a slice of type string
func ReadFile(filePath string) []string {
    readFile, err := os.Open(filePath)

    if err != nil {
    }

    fileScanner := bufio.NewScanner(readFile)

    fileScanner.Split(bufio.ScanLines)

    var output []string

    for fileScanner.Scan() {
        output = append(output, fileScanner.Text())
    }

    return output
}


func StringSliceToIntegerSlice(lines []string) ([]int, error) {
    integers := make([]int, 0, len(lines))

    for _, line := range lines {
        n, err := strconv.Atoi(line)
        if err != nil {
            return nil, err
        }
        integers = append(integers, n)
    }
    return integers, nil
}


func Regexp2FindAllString(re *regexp2.Regexp, s string) []string {
	var matches []string
	m, _ := re.FindStringMatch(s)
	for m != nil {
		matches = append(matches, m.String())
		m, _ = re.FindNextMatch(m)
	}
	return matches
}
