package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func problemPuller(fileName string) ([]problem, error) {

	//1.Open the file
	if fObj, err := os.Open(fileName); err == nil {

		//2.Create a new reader
		csvR := csv.NewReader(fObj)

		//3.Read the files
		if cLines, err := csvR.ReadAll(); err == nil {
			//4.Call the parseProblem function
			return parseProblem(cLines), nil
		} else {
			return nil,
				fmt.Errorf("error in readinf date in csv"+
					"format from %s file; %s", fileName, err.Error())
		}

	} else {
		return nil,
			fmt.Errorf("error in opening %s file; %s", fileName, err.Error())
	}
}

func parseProblem(lines [][]string) []problem {

	r := make([]problem, len(lines))

	for i:=0;i<len(lines);i++{
		r[i] = problem{question: lines[i][0], answer:lines[i][1]}
	}

	return r
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}

type problem struct {
	question string
	answer   string
}

func main() {
	//1. Input the name of the file
	fName := flag.String("f", "quiz.csv", "path for the csv file")

	//2. Set the duration of the timer
	timer := flag.Int("t", 30, "timer for the quiz")
	flag.Parse()

	//3. Pull the problems from the file (calling the problem parser function)
	problems, err := problemPuller((*fName))

	//4. Handle the error
	if err != nil {
		exit(fmt.Sprintf("something went wrong:%s", err.Error()))
	}

	//5. Create a variable to count the correct answers
	correctAns := 0

	//6. Using the duration of timer, we want to initialize the timer
	tObj := time.NewTimer(time.Duration(*timer) * time.Second)
	ansC := make(chan string)

	//7. Loop through the problems, print the questions, we will accept the answers
	problemLoop:

		for i,p := range problems{
			var answer string

			fmt.Printf("Problem %d : %s", i+1, p.question)

			go func(){
				fmt.Scanf("%s", &answer)
				ansC<- answer
			}()

			select{
				case <- tObj.C:
					fmt.Println()
					break problemLoop
				
				case iAns := <-ansC:
					if iAns == p.answer{
						correctAns++
					}
					if i== len(problems){
						close(ansC)
					}
			}

		}

	//8. We will calculate and print the result
	fmt.Printf("Your result is %d out of %d/n", correctAns, len(problems))
	fmt.Printf("Press enter to exit")
	<-ansC
}
