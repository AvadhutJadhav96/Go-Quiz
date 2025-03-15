package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func problemPuller(fileName string) ([]problem, error) {
	// Open the file
	fObj, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error in opening %s file: %s", fileName, err.Error())
	}
	defer fObj.Close() // Ensure the file is closed

	// Create a new CSV reader
	csvR := csv.NewReader(fObj)

	// Read the CSV file
	cLines, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error in reading data from %s file: %s", fileName, err.Error())
	}

	// Parse the problems
	return parseProblem(cLines), nil
}

func parseProblem(lines [][]string) []problem {
	r := make([]problem, len(lines))

	for i := 0; i < len(lines); i++ {
		r[i] = problem{question: lines[i][0], answer: lines[i][1]}
	}

	return r
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

type problem struct {
	question string
	answer   string
}

func main() {
	// Input the name of the file
	fName := flag.String("f", "quiz.csv", "path for the csv file")

	// Set the duration of the timer
	timer := flag.Int("t", 30, "timer for the quiz")
	flag.Parse()

	// Pull the problems from the file
	problems, err := problemPuller(*fName)
	if err != nil {
		exit(fmt.Sprintf("Something went wrong: %s", err.Error()))
	}

	// Variable to count correct answers
	correctAns := 0

	// Initialize the timer
	tObj := time.NewTimer(time.Duration(*timer) * time.Second)

	// Loop through the problems
	for i, p := range problems {
		fmt.Printf("Problem %d: %s ", i+1, p.question)

		ansC := make(chan string) // Create a new channel for each question

		// Take input in a separate goroutine
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer) // Read user input
			ansC <- answer
		}()

		select {
		case <-tObj.C:
			fmt.Println("\nTime's up!")
			goto END // Exit the loop when the timer expires
		case iAns := <-ansC:
			if iAns == p.answer {
				correctAns++
			}
		}
	}

END:
	// Print the result
	fmt.Printf("\nYour result is %d out of %d\n", correctAns, len(problems))

	// Wait for user to press enter before exiting
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
