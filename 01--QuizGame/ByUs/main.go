package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type QuizUnit struct {
	question string
	answer   string
}

func main() {
	quizUnits := readQuizUnits()
	fmt.Printf("Quiz Units: %v", quizUnits)
}

func readQuizUnits() []QuizUnit {
	data, err := os.Open("./problems.csv")
	if err != nil {
		log.Fatalf("Error reading file:\n%s", err)
	}

	reader := csv.NewReader(data)

	quizUnits := []QuizUnit{}
	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Error reading line in file:\n%s", err)
		}

		quizUnits = append(quizUnits, QuizUnit{
			question: line[0],
			answer:   line[1],
		})
	}

	return quizUnits
}
