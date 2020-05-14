package main

import (
	"syscall"
	"os/signal"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/docopt/docopt-go"
)

const (
	version          = "1.0.0"
	defaultTimeLimit = 90.0
)

const usage string = `Go Quiz.

Usage:
	go_quiz run <filename> [--timelimit=<tl>]
	go_quiz -h | --help
	go_quiz --version

Arguments:
	filename      Filename/path of csv file format containing quiz problem-answer

Options:
	-h --help          Show this screen.
	--version          Show version.
	--timelimit=<tl>   Set time Limit for quiz [default: 30.0].`

var config struct {
	Run       bool
	TimeLimit float64 `docopt:"--timelimit"`
	Filename  string  `docopt:"<filename>"`
	Version   bool    // flag --version
}

// Question represent problem question which holds data to
// problem statement and ground truth answer
type Question struct {
	Statement string
	Answer    string
}

// QuizSession represent quiz session which holds data to
// io buffer, score, timelimit (sec)
type QuizSession struct {
	Questions       []Question
	TimeLimit       float64
	Score           int
	LastQuestionIdx int
}

func main() {
	// Parse CLI args & options
	opts, _ := docopt.ParseDoc(usage)
	opts.Bind(&config)

	// >_ go_quiz --version
	if config.Version == true {
		fmt.Println(version)
		return
	}

	var sigIntrChan = make(chan os.Signal)
	signal.Notify(sigIntrChan, os.Interrupt, syscall.SIGINT)

	// >_ go_quiz run <filename>
	if config.Run == true {
		quizSess, err := NewQuizSession(config.Filename, config.TimeLimit)
		if err != nil {
			return
		}

		// handle signal interupt, print summary then exit program 
		go func() {
			<-sigIntrChan
			printSummary(quizSess)
			os.Exit(0)
		}()

		runQuiz(quizSess)
		printSummary(quizSess)
	}
}

// NewQuizSession create new QuizSession with initalized information
func NewQuizSession(filename string, timelimit float64) (*QuizSession, error) {
	var quizSess *QuizSession

	csvFile, err := os.OpenFile(filename, os.O_RDONLY, 0755)
	if err != nil {
		log.Fatalln("Unable to open quiz file: ", filename)
		return quizSess, err
	}

	csvReader := csv.NewReader(csvFile)

	if timelimit < 0 {
		timelimit = defaultTimeLimit
	}

	questions, err := parseQuestions(csvReader)
	if err != nil {
		log.Fatalln("Unable to parse file as CSV: ", filename)
		return quizSess, err
	}

	quizSess = &QuizSession{
		Questions:       questions,
		TimeLimit:       timelimit,
		LastQuestionIdx: 0,
		Score:           0,
	}
	return quizSess, nil
}

func parseQuestions(csvReader *csv.Reader) ([]Question, error) {
	var questions []Question

	records, err := csvReader.ReadAll()
	if err != nil {
		return questions, err
	}

	for _, record := range records {
		question := Question{
			Statement: strings.TrimSpace(record[0]),
			Answer:    strings.TrimSpace(record[1]),
		}
		questions = append(questions, question)
	}
	return questions, nil
}

func runQuiz(quizSess *QuizSession) (*QuizSession, error) {
	questionFormat := "# Question %02d: %v\n"
	timer := time.NewTimer(time.Duration(quizSess.TimeLimit) * time.Second)

problemloop:
	for i, question := range quizSess.Questions {
		var userChan = make(chan string)

		statement, answer := question.Statement, question.Answer
		fmt.Printf(questionFormat, i, statement)

		go func() {
			var userAnswer string
			fmt.Print("Answer: ")
			fmt.Scanf("%s\n", &userAnswer)
			userChan <- userAnswer
		}()

		select {
		case <-timer.C:
			break problemloop
		case userAnswer := <-userChan:
			if answer == userAnswer {
				quizSess.Score++
			}
		}
		quizSess.LastQuestionIdx = i
	}
	return quizSess, nil
}

func printSummary(quizSess *QuizSession) {
	numQuestions := len(quizSess.Questions)

	fmt.Println()

	// Print status: Completed / Time's Up
	if quizSess.LastQuestionIdx < len(quizSess.Questions)-1 {
		fmt.Println("================= Time's Up ! =================")
	} else {
		fmt.Println("================= Completed ! =================")
	}

	// Print simple statistics
	fmt.Printf(".. Your Score: %02d / %02d\n", quizSess.Score, numQuestions)
	fmt.Printf(".. # of answered questions: %02d\n", quizSess.LastQuestionIdx)
	fmt.Printf(".. # of wrong answer: %02d\n", quizSess.Score)
	fmt.Printf(".. # of missed question %02d\n", numQuestions-quizSess.LastQuestionIdx)
}
