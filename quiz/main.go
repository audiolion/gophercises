package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))

	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return problems
}

func main() {
	var (
		csvFilename = flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
		timeLimit   = flag.Int("limit", 30, "the time limit for the quiz in seconds")
	)
	flag.Parse()

	file, err := os.Open(*csvFilename)
	defer file.Close()
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

	doneCh := make(chan bool)

	correct := 0

	go func() {
		for i, p := range problems {
			msg := "Problem #%d: %s = \n"
			fmt.Printf(msg, i+1, p.question)
			var answer string
			fmt.Scanf("%s\n", &answer)
			if answer == p.answer {
				correct++
			}
			if i == len(problems)-1 {
				doneCh <- true
			}
		}
	}()

	go func() {
		timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
		<-timer.C
		fmt.Printf("\n\nYour time is up.")
		doneCh <- true
	}()

	<-doneCh
	fmt.Printf("\nYou scored %d out of %d\n", correct, len(problems))

}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
