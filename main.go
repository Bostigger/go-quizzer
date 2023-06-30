package main

import (
	flag2 "flag"
	"fmt"
	"github.com/bostigger/go-quizzer/controller"
)

func main() {
	fmt.Println("GO QUIZZER")
	file := flag2.String("f", "quiz.csv", "Filename")
	flag2.Parse()
	questions, err := controller.ProblemPuller(*file)
	if err != nil {
		fmt.Println("Something went wrong pulling from csv file", err)
		return
	}
	ansChan := make(chan string)

	controller.ProblemLooper(questions, ansChan)
}
