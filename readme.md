# Quiz Application

## Overview
This is a command-line quiz application written in **Go**. The program reads questions from a CSV file, asks the user each question, and waits for an answer. The quiz runs with a timer, and once time is up, the quiz ends, displaying the final score.

## Features
- Reads quiz questions from a CSV file.
- Supports a configurable timer to limit the quiz duration.
- Uses concurrency (goroutines) for handling user input.
- Displays the final score at the end of the quiz.

## Prerequisites
- **Go** installed on your system ([Download Go](https://golang.org/dl/)).
- A CSV file containing the quiz questions (default: `quiz.csv`).

## Installation
1. Clone this repository or download the source code:
   ```sh
   git clone https://github.com/yourusername/quiz-app.git
   cd quiz-app
   ```
2. Build the project:
   ```sh
   go build -o quiz
   ```

## Usage
Run the program with default settings:
```sh
./quiz
```

Specify a different CSV file and timer duration:
```sh
./quiz -f custom_quiz.csv -t 60
```

## CSV File Format
The quiz questions are stored in a CSV file with two columns:
```
question,answer
5+5,10
What is the capital of France?,Paris
```
Ensure the file does **not** have a header row.

## How It Works
1. The program reads the CSV file and loads the questions.
2. It starts a timer for the specified duration (default: 30 seconds).
3. Questions are displayed one by one, and user input is taken.
4. If the user provides the correct answer, the score increases.
5. If the timer runs out, the quiz ends immediately.
6. The final score is displayed at the end.

## Example Run
```
Problem 1: 5+5
> 10
Problem 2: What is the capital of France?
> Paris

Your result is 2 out of 2
Press Enter to exit...
```



