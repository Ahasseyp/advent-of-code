package main

import (
	"fmt"
	"math"
	"regexp"

	//"sort"
	//"strconv"
	//"strings"

	"advent-of-code/utils"
)

var INITIAL_NODE = "AAA"
var FINAL_NODE = "ZZZ"

func main() {
	text := utils.ReadFile("./input.txt")

    instructions := text[0]

    text = text[2:]

    var network = make(map[string]map[string]string)

	for _, textLine := range text {
        wordExpressions := regexp.MustCompile(`(\w+)[^\w]+(\w+)[^\w]+(\w+)`)
        matches := wordExpressions.FindAllStringSubmatch(textLine, -1)

        key := matches[0][1]
        left := matches[0][2]
        right := matches[0][3]

        network[key] = map[string]string{"L": left, "R": right}
	}

    currentNode := INITIAL_NODE

    var totalSteps int

    for i := 0; currentNode != FINAL_NODE; i++ {
        step := math.Mod(float64(i), float64(len(instructions)))
        instruction := string(instructions[int(step)])
        currentNode = network[currentNode][instruction]
        totalSteps = i + 1
    }

    fmt.Println(totalSteps)
}
