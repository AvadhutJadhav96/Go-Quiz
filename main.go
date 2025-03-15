package main

import(
	"flag"
)

func problemPuller(fileName string)([]problem, error){

}

func parseProblem(lines [][]string) []problem{
	
}

type problem struct{
	question string
	answer   string
}

func main(){
	//1. Input the name of the file
	fName := flag.String("f","quiz.csv", "path for the csv file")

	//2. Set the duration of the timer
	timer := flag.Int("t", 30, "timer for the quiz")
	flag.Parse()

	//3. Pull the problems from the file (calling the problem parser function)
	problems,err := problemPuller((*fName))

	//4. Handle the error
	if err!=nil{
		exit(fmt.Sprintf("something went wrong:%s", err.Error()))
	}

	//5. Create a variable to count the correct answers
	correctAns :=0

	//6. Using the duration of timer, we want to initialize the timer
	tObj :=time.NewTimer(time.Duration(*timer)*time.Second)
	ansC :=make(chan string)
	
	//7. Loop through the problems, print the questions, we will accept the answers
	//8. We will calculate and print the result
}