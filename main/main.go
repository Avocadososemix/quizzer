package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Problem struct {
	question string
	answer   string
}

func readCSV() []Problem {
	csvFile, err := os.Open("data/problems.csv")
	check(err)

	reader := csv.NewReader(bufio.NewReader(csvFile))
	var problems []Problem

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		problems = append(problems, Problem{
			question: line[0],
			answer:   line[1],
		})
	}
	return problems
}

func main() {

	var problems = readCSV()
	var correctCount, questionCount Counter = new(counter8), new(counter8)
	reader := bufio.NewReader(os.Stdin)

	timeLimit := time.Duration(30)
	fmt.Printf("Starting quiz. Stop by typing x. Time limit is %d seconds.\n", timeLimit)
	timer := time.NewTimer(time.Second * timeLimit)
	defer timer.Stop()
	go func() {
		<-timer.C
		fmt.Printf("Time's up! Correct answers: %s, total questions answered: %s\n", correctCount.Value(), questionCount.Value() )
		os.Exit(0)
	}()

	for i :=0;i<100 ;i++  {
		fmt.Printf("What is the answer to: %s\n", problems[i].Question())
		text, _ := reader.ReadString('\n')
		if strings.TrimSpace(text) == "x" {
			fmt.Println("Quitting.")
		break
		} else if (strings.TrimSpace(text) == problems[i].Answer()) {
			fmt.Println("That's correct!")
			correctCount.Inc()
			questionCount.Inc()
		} else {
			questionCount.Inc()
			fmt.Printf("Sorry, the correct answer is: %s\n", problems[i].Answer())
		}
	}
	fmt.Printf("Correct answers: %s, total questions answered: %s\n", correctCount.Value(), questionCount.Value() )
}

func (f *Problem) Question() string {
	return f.question
}

func (f *Problem) Answer() string {
	return f.answer
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

type Counter interface {
	Inc()
	Value() string
}

type counter8 uint8

func (c *counter8) Inc()      { *c++ }
func (c counter8) Value() string { return strconv.FormatUint(uint64(c),10) }


//https://github.com/gophercises/quiz#part-2