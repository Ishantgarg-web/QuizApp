package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

// define structs
type problems struct {
	Category         string   `json:"category"`
	Id               string   `json:"id"`
	CorrectAnswer    string   `json:"correctanswer"`
	IncorrectAnswers []string `json:"incorrectanswers"`
	Question         string   `json:"question"`
	Difficulty       string   `json:"difficulty"`
}

func query(category, difficulty int) ([]problems, error) {
	var s string
	if category == 1 {
		s = "arts_and_literature"
	} else if category == 2 {
		s = "film_and_tv"
	} else if category == 3 {
		s = "food_and_drink"
	} else if category == 4 {
		s = "general_knowledge"
	} else if category == 5 {
		s = "geography"
	} else if category == 6 {
		s = "history"
	} else if category == 7 {
		s = "music"
	} else if category == 8 {
		s = "science"
	} else if category == 9 {
		s = "society_and_culture"
	} else if category == 10 {
		s = "sport_and_leisure"
	}
	var diff string
	if difficulty == 1 {
		diff = "easy"
	} else if difficulty == 2 {
		diff = "medium"
	} else {
		diff = "hard"
	}
	var url = "https://the-trivia-api.com/api/questions?categories=" + s + "&limit=10&difficulty=" + diff
	resp, err := http.Get(url)
	if err != nil {
		return []problems{}, err
	}
	defer resp.Body.Close()
	var p []problems
	json.NewDecoder(resp.Body).Decode(&p)
	return p, nil
}

// showproblems -> Here we will show problems based on user Category and Difficulty

func showProblems(data []problems) int {
	correctAnswer := 0
	for i := 0; i < len(data); i++ {
		d := data[i]
		fmt.Println("Question ", i+1, " is:")
		fmt.Println(d.Question)
		fmt.Println("Options are: ")

		// make a slice where we can store all options
		var options []string
		options = append(options, d.CorrectAnswer)
		//fmt.Println("Incorrect: ", d.IncorrectAnswers[0])
		options = append(options, d.IncorrectAnswers[0])
		options = append(options, d.IncorrectAnswers[1])
		options = append(options, d.IncorrectAnswers[2])
		rand.Shuffle(len(options), func(i, j int) { options[i], options[j] = options[j], options[i] })
		fmt.Println("1. " + options[0] + "\n2. " + options[1] + "\n3. " + options[2] + "\n4. " + options[3])
		fmt.Println("Type your answer (Please type corressponding number")
		var userAnswer int
		fmt.Scanln(&userAnswer) //taking user answer
		if options[userAnswer-1] == d.CorrectAnswer {
			fmt.Println("Correct Answer")
			correctAnswer++
		} else {
			fmt.Println("InCorrect Answer")
		}
	}
	return correctAnswer
}

func main() {

	// define options for category

	fmt.Println("#####   Welcome to Quiz Project ######\nHere you can test your knowledge....")
	fmt.Println("Please choose Valid category( Type Corresponding number)")
	fmt.Println("1. Arts & Literature\n2. Film & TV\n3. Food & Drink\n4. General Knowledge\n5. Geography\n6. History\n7. Music\n8. Science\n9. Society and Culture\n10. Sport and Leisure")
	var category int
	fmt.Scanln(&category)

	// define options for Difficulty
	fmt.Println("Please choose Difficulty Medium( Type Corresponding number)")
	fmt.Println("1. Easy\n2. Medium\n3. Hard")
	var difficulty int
	fmt.Scanln(&difficulty)

	data, err := query(category, difficulty)
	if err != nil {
		log.Fatal(err)
	}
	correctAnswer := showProblems(data)
	fmt.Println("Test is Completed!!")
	fmt.Println("Your scoreCard:\nCorrect Answers are: ", correctAnswer, "\nIncorrect Answers are: ", 10-correctAnswer)
	fmt.Println("\n\n Hope you like the Quiz App.")
}
