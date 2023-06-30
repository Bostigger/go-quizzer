package controller

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/bostigger/go-quizzer/model"
	"os"
	"strings"
	"time"
)

func ProblemPuller(fileName string) ([]model.QuizQuestion, error) {
	data, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("error opening file %s", err.Error())
		return nil, err
	}
	csvR := csv.NewReader(data)
	csvLines, err := csvR.ReadAll()
	if err != nil {
		fmt.Printf("error reading file %s", err.Error())
		return nil, err
	}
	questions := problemParser(csvLines)
	if err != nil {
		fmt.Printf("error parsing file %s", err.Error())
		return nil, err
	}
	return questions, nil
}

func problemParser(lines [][]string) []model.QuizQuestion {
	r := make([]model.QuizQuestion, len(lines))
	for i := 0; i < len(lines); i++ {
		r[i] = model.QuizQuestion{
			Question: lines[i][0],
			Answer:   lines[i][1],
		}
	}
	return r
}
func ProblemLooper(questions []model.QuizQuestion, answerCh chan string) {
	correctAnswer := 0
	go func() {
		for i, question := range questions {
			fmt.Printf("Problem %d : %s=", i+1, question.Question)
			reader := bufio.NewReader(os.Stdin)
			answer, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("error reading user answer %s", err.Error())
				return
			}
			answer = strings.TrimSpace(answer)
			if answer == question.Answer {
				correctAnswer++
			}
		}
		close(answerCh)
	}()
	select {
	case <-answerCh:
		fmt.Printf("Your score is %d out of %d\n", correctAnswer, len(questions))
	case <-time.After(time.Second * 30):
		fmt.Printf("\n Time's up!!!")
		fmt.Printf("Your score is %d out of %d\n", correctAnswer, len(questions))
	}

}
