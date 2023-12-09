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
	data, err := os.Open("./problems.csv")
	if err != nil {
		log.Fatalf("Error reading file:\n%s", err)
		return
	}

	reader := csv.NewReader(data)

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Error reading line in file:\n%s", err)
			return
		}

		quizUnit := QuizUnit{
			question: line[0],
			answer:   line[1],
		}

		fmt.Printf("%v", quizUnit)
	}

}
