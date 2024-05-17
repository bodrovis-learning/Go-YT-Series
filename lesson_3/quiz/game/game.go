package game

import (
	"bufio"
	"fmt"
	"os"
	"quiz/questions"
	"strings"
)

func Run(questions []questions.Question) (correctAnswers uint) {
	fmt.Println("Welcome to the Country Quiz!")

	for _, question := range questions {
		if askQuestion(question) {
			correctAnswers++
		}
	}

	return correctAnswers
}

func askQuestion(question questions.Question) bool {
	fmt.Printf("\nEnter the capital of %s: ", question.Country)

	if getUserInput() == strings.ToLower(question.Capital) {
		fmt.Println("Correct!")
		return true
	} else {
		fmt.Printf("Incorrect! The correct answer is %s.\n", question.Capital)
		return false
	}
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Your answer: ")
		result, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred while reading the entered text! Please try again...")
			continue
		}

		return strings.ToLower(strings.TrimRight(result, "\r\n"))
	}
}
