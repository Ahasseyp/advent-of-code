package main

import (
    "fmt"
    "regexp"
    //"strings"
	//"math"
    //"slices"
	//"strconv"

	"advent-of-code/utils"
)

type Race struct {
    time int
    distance int
}

func distanceTravelled(holdTime int, raceTime int) int {
    return holdTime * (raceTime - holdTime)
}

func (race Race) waysToWin() int {
    waysToWin := 0

    for time := 0; time < race.time; time++ {
        if distanceTravelled(time, race.time) > race.distance {
            waysToWin += 1
        }
    }

    return waysToWin
}

func main()  {

    text := utils.ReadFile("./input.txt")

    numberExpression := regexp.MustCompile(`\d+`)

    timesText := numberExpression.FindAllString(text[0], -1)
    distancesText := numberExpression.FindAllString(text[1], -1)

    times, _ := utils.StringSliceToIntegerSlice(timesText)
    distances, _ := utils.StringSliceToIntegerSlice(distancesText)

    var races []Race

    for i := 0; i < len(times); i++ {
        races = append(races, Race{time: times[i], distance: distances[i]})
    }

    result := 1

    for _, race := range races {
        result *= race.waysToWin()
    }

    fmt.Println(result)
}
