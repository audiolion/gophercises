package main

/*
 * Eight houses, represented as cells, are arranged in a straight line.
 * Each day every cell competes with its adjacent cells (neighbors).
 * An integer value of 1 represents an active cell and a value of 0 represents
 * an inactive cell. If the neighbors on both the sides of a cell are either
 * active or inactive, the cell becomes inactive on the next day; otherwise
 * the cell becomes active. The two cells on each end have a single adjacent
 * cell, so assume that the unoccupied space on the opposite side is an
 * inactive cell. Even after updating the cell state, consider its previous
 * state when updating the state of other cells. The state information of all
 * cells should be updated simultaneously.
 *
 * Write an algorithm to output the state of the cells after the given number of days.
 *  Input :
 *    1 0 0 0 0 1 0 0
 *    1
 *  Output :
 *    0 1 0 0 1 0 1 0
 */

import (
	"flag"
	"fmt"
)

const (
	inactive byte = byte(0)
	active   byte = byte(1)
)

func main() {
	cells := flag.String("cells", "", "eight 0s or 1s")
	days := flag.Int("days", 0, "number of days to run the simulation")
	flag.Parse()

	houses := make([]byte, 10, 10)
	for i, ch := range *cells {
		if string(ch) == "0" {
			houses[i+1] = inactive
		} else {
			houses[i+1] = active
		}
	}

	for i := 0; i < *days; i++ {
		nextState := make([]byte, 10, 10)
		for j := 1; j < len(houses)-1; j++ {
			nextState[j] = compareNeighbors(j, houses)
		}
		houses = nextState
	}

	fmt.Println(houses[1:9])
}

func compareNeighbors(i int, houses []byte) (nextState byte) {
	if houses[i-1] == houses[i+1] {
		nextState = inactive
	} else {
		nextState = active
	}
	return
}
