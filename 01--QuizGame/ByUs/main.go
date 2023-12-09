package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/debug"
)

type QuizUnit struct {
	question string
	answer   string
}

func main() {
	filename := flag.String("f", "problems.csv", "path to csv quiz data")
	flag.Parse()

	quizUnits := readQuizUnits(*filename)
	fmt.Printf("Quiz Units: %v", quizUnits)
}

func readQuizUnits(filename string) []QuizUnit {
	data, err := os.Open(filename)

	if err != nil {
		errorMessage := fmt.Sprintf("Error reading file: %s\n%s", err.Error(), debug.Stack())
		log.Output(2, errorMessage)
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
