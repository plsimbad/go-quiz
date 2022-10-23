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
	timeLimit := flag.Int("t", 10, "a int")
	flag.Parse()

	qLst, err := readCSVFile(*path)

	if err != nil {
		fmt.Println(err)
		return
	}

	correct := askUser(qLst, *timeLimit)

	fmt.Println("answered", correct, "of", len(qLst), "questions")
}

func askUser(qLst []questions, timeLimit int) int {
	correct := 0
	answer := make(chan string)
	go getInput(answer)

	for _, q := range qLst {
		timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

		res, err := askQ(q, timer.C, answer)
		if err != nil {
			fmt.Println(err)
			return correct
		}

		correct += res
	}

	return correct
}

func getInput(answer chan string) {
	for {
		in := ""
		fmt.Scanln(&in)
		answer <- in
	}
}

func askQ(question questions, timer <-chan time.Time, answer <-chan string) (int, error) {
	fmt.Println(question.question)
	for {
		select {
		case <-timer:
			return 0, fmt.Errorf("Time out")
		case ans := <-answer:
			if question.answer == ans {
				return 1, nil
			}

			return 0, nil
		}

	}
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
