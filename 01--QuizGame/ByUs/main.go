package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

type QuizUnit struct {
	question string
	answer   string
}

func main() {
	filename := flag.String("f", "problems.csv", "path to csv quiz data")
	timePerQuestion := flag.Int("t", 5, "time per question, in seconds")

	flag.Parse()

	quizUnits := readQuizUnits(*filename)
	score := 0

	for i, qunit := range quizUnits {
		responseChan := make(chan string, 1)

		go askQuestion(i, qunit.question, responseChan)

		select {
		case userResponse := <-responseChan:
			if userResponse != "" && strings.TrimSpace(userResponse) == qunit.answer {
				score++
			}

		case <-time.After(time.Duration(*timePerQuestion) * time.Second):
			fmt.Println("\nTime elapsed for that question!")
		}
	}

	fmt.Print("\nCongratulations on completing the quiz!\n")
	fmt.Print("Your score is: ", score, " out of ", len(quizUnits))
}

func askQuestion(questionNumber int, question string, c chan string) {
	userResponse := ""

	fmt.Printf("\nQ%d: %s = ", questionNumber, question)
	fmt.Scanln(&userResponse)

	c <- userResponse
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

		if line[0] == "" || line[1] == "" {
			continue
		}

		quizUnits = append(quizUnits, QuizUnit{
			question: line[0],
			answer:   line[1],
		})
	}

	return quizUnits
}
