package utils

import (
    "bufio"
    "os"
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
