# Gophercises: Exercise 1 - Quiz Game

[![exercise status: released](https://img.shields.io/badge/exercise%20status-released-green.svg?style=for-the-badge)](https://gophercises.com/exercises/quiz)

## Documentations

* Overview (This README.md)
* More on exercise detail: [gophercise](docs/EXERCISE.md)


## Quiz Game
In this exercise, we build a CLI Quiz Game. The user can interactively answer question in the command prompt. The questions are suppplied through a CSV file with the correct answer as well. 

The programs run as follows:

```
# Question 01: 1+1
Answer: 2
# Question 02: 8+3
Answer: 11
# Question 03: 1+2
Answer: 4
# Question 04: 8+6
Answer: 14
# Question 05: 3+1
Answer: 4
# Question 06: 1+4
Answer: 5
# Question 07: 5+1
Answer: 6
# Question 08: 2+3
Answer: 
...

================= Time's Up ! =================
Your Score: 07 / 12
# of answered questions: 07
# of wrong answer: 07
# of missed question 05
```

## Requirements
* Golang Compiler > 1.7.0

## Setup and User guide 

### Usage

``` bash
Go Quiz.

Usage:
	go_quiz run <filename> [--timelimit=<tl>]
	go_quiz -h | --help
	go_quiz --version

Arguments:
	filename      Filename/path of csv file format containing quiz problem-answer

Options:
	-h --help          Show this screen.
	--version          Show version.
	--timelimit=<tl>   Set time Limit for quiz [default: 30.0].
```

### CSV File Input

Prepare CSV file for quiz. The file contains two columns: `<problem_statement>, <answer>`. 

For example,
```csv
10+4,14
90/3,30
who is Jeff?, Bezos
my age?, 22
...
```
The sample CSV file has been provided in this repository `problem.csv`

### Setup, Build, & Run
1. Install golang dependency
```
# one dependency only
$ go get github.com/docopt/docopt-go
```

2. First build the program using golang
```
$ go build -o go_quiz *.go
```

3. Run the program as follows:
```
$ go_quiz run problem.csv --timelimit=30.0
```

## TODO - Open Tasks Left
See [EXERCISE.md](docs/EXERCISE.md) for more details on the exercise.
- [x] Part I  - Simple version, no time limit
- [x] Part II - Adding Concurrency channel for input scan and timer
- [x] Bonus 1 - Trimmed and cleanup string input for white prefix and trailing spaces.
- [ ] Bonus 2 - Shuffle quiz questions

## Reference
* Gophercise - Exercises for budding gophers (Go developers) by [Jon Calhoun](https://github.com/gophercises)
* Go Docopts - Parsing CLI args and opts [docopts-go](https://github.com/docopt/docopt.go) 
