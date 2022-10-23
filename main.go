package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type questions struct {
	question string
	answer   string
}

func main() {

	path := flag.String("f", "./problems.csv", "a string")
	// timeInSec := flag.Int("t", 10, "a int")
	flag.Parse()

	qLst, err := readCSVFile(*path)

	if err != nil {
		fmt.Println(err)
		return
	}

	correct := 0

	for _, q := range qLst {
		fmt.Println(q.question)
		answer := ""
		fmt.Scanln(&answer)
		if answer == q.answer {
			correct++
		}
	}

	fmt.Println("answered", correct, "of", len(qLst), "questions")
}

func readCSVFile(path string) ([]questions, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	lines, err := csv.NewReader(f).ReadAll()

	retLst := make([]questions, 0)

	for _, line := range lines {
		retLst = append(retLst, questions{question: line[0], answer: line[1]})
	}

	return retLst, nil
}

func startTimer(timeInSec int) *time.Timer {
	return time.NewTimer(time.Duration(timeInSec) * time.Second)
}
