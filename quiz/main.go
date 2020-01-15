package main

import (
	"flag"
	"fmt"
	"os"
	"encoding/csv"
	"strings"
)

type Problem struct {
	question string
	answer string
}

func parseLines(lines [][]string) []Problem {
	problems := make([]Problem, len(lines))

	for i, line := range lines {
		problems[i] = Problem{
			question: line[0],
			answer: strings.TrimSpace(line[1]),
		}
	}

	return problems
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		msg := "Failed to open the CSV file: %s\n"
		exit(fmt.Sprintf(msg, *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		msg := "Failed to parse provided CSV file"
		exit(msg)
	}

	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		msg := "Problem #%d: %s = \n"
		fmt.Printf(msg, i+1, p.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			correct++
		}
	}

	msg := "You scored %d out of %d\n"
	fmt.Printf(msg, correct, len(problems))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
