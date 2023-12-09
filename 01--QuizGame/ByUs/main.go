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
	quizUnits := readQuizUnits("problemss.csv")
	fmt.Printf("Quiz Units: %v", quizUnits)
}

func readQuizUnits(filename string) []QuizUnit {
	data, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Error reading file:\n%s", err)
	}

	defer data.Close()

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
