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
	flag.Parse()

	quizUnits := readQuizUnits(*filename)
	score := 0

	for i, qunit := range quizUnits {
		responseChan := make(chan string, 1)
		timerChan := make(chan struct{}, 1)

		go askQuestion(i, qunit.question, responseChan)
		go timer(5, timerChan)

		select {
		case userResponse := <-responseChan:
			fmt.Println("you entered: ", userResponse)
			if userResponse != "" && strings.TrimSpace(userResponse) == qunit.answer {
				score++
			}
			continue

		case <-timerChan:
			continue
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

func timer(seconds int, c chan struct{}) {
	time.Sleep(time.Duration(seconds) * time.Second)
	c <- struct{}{}
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
