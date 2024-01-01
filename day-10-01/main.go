package main

import (
	"fmt"
	"math"
	"slices"

	"advent-of-code/utils"
)

type Pipe struct {
	x, y                     int
	north, east, south, west *Pipe
	typeOf                   string
}
func (pipe Pipe) String() string {
    return fmt.Sprintf("<Pipe {%v %v} %v>", pipe.x, pipe.y, pipe.typeOf)
}
func (pipe Pipe) availableConnections() []Pipe {
    var availableConnections []Pipe

    if pipe.north != nil {
        availableConnections = append(availableConnections, *pipe.north)
    }
    if pipe.east != nil {
        availableConnections = append(availableConnections, *pipe.east)
    }
    if pipe.south != nil {
        availableConnections = append(availableConnections, *pipe.south)
    }
    if pipe.west != nil {
        availableConnections = append(availableConnections, *pipe.west)
    }

    return availableConnections
}

func createNetwork(text []string) [][]Pipe {
    // Create 2D slice network
	var network = make([][]Pipe, len(text))
	for v := range network {
		network[v] = make([]Pipe, len(text[v]))
	}

    // Set Pipe on every coordinate of the 2D slice
	for y, textLine := range text {
		for x, char := range textLine {
			typeOf := string(char)
			network[y][x] = Pipe{x: x, y: y, typeOf: typeOf}
		}
	}

	return network
}

func setNetworkRelations(network [][]Pipe) (newNetwork [][]Pipe, startingPipe Pipe) {
    for y := 1; y < (len(network) - 1); y++ {
        row := network[y]
        for x := 1; x < (len(row) - 1); x++ {
            pipe := &network[y][x]
            pipeNorth := &network[y-1][x]
            pipeSouth := &network[y+1][x]
            pipeEast := &network[y][x+1]
            pipeWest := &network[y][x-1]
            switch pipe.typeOf {
            case "S":
                startingPipe = *pipe
                valuesNorth := []string{"|", "F", "7"}
                valuesSouth := []string{"|", "J", "L"}
                valuesEast := []string{"-", "J", "7"}
                valuesWest := []string{"-", "F", "L"}

                if slices.Contains(valuesNorth, pipeNorth.typeOf) {
                    network[y][x].north = pipeNorth
                }
                if slices.Contains(valuesSouth, pipeSouth.typeOf) {
                    network[y][x].south = pipeSouth
                }
                if slices.Contains(valuesEast, pipeEast.typeOf) {
                    network[y][x].east = pipeEast
                }
                if slices.Contains(valuesWest, pipeWest.typeOf) {
                    network[y][x].west = pipeWest
                }
            case "|":
                network[y][x].north = pipeNorth
                network[y][x].south = pipeSouth
            case "-":
                network[y][x].west = pipeWest
                network[y][x].east = pipeEast
            case "L":
                network[y][x].north = pipeNorth
                network[y][x].east = pipeEast
            case "J":
                network[y][x].north = pipeNorth
                network[y][x].west = pipeWest
            case "7":
                network[y][x].south = pipeSouth
                network[y][x].west = pipeWest
            case "F":
                network[y][x].south = pipeSouth
                network[y][x].east = pipeEast
            default:
                continue
            }
		}
	}

	return network, network[startingPipe.y][startingPipe.x]
}

func main() {
	text := utils.ReadFile("./input.txt")

	text = utils.PadSlice(text, ".")

	network := createNetwork(text)

    network, startingPipe := setNetworkRelations(network)

    var visitedPipes = []Pipe{startingPipe}

    nextPipe := startingPipe.availableConnections()[0]

    for nextPipe != startingPipe{
        visitedPipes = append(visitedPipes, nextPipe)
        for _, availablePipe := range nextPipe.availableConnections() {
            if slices.Contains(visitedPipes, availablePipe) {
                continue
            } else {
                nextPipe = availablePipe
            }
        }
        if nextPipe == visitedPipes[len(visitedPipes) - 1] {
            break
        }
    }

    visitedPipes = visitedPipes[1:]

    fmt.Println(math.Ceil(float64(len(visitedPipes))/2))
}
