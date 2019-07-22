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

func main() {
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

	reader2 := bufio.NewReader(os.Stdin)

	var counter, questionCount Counter
	counter = new(counter8)
	questionCount = new(counter8)
	fmt.Println("Starting quiz. Stop by typing x. Time limit is 30 seconds.")
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()
	go func() {
		<-timer.C
		fmt.Printf("Time's up! Correct answers: %s, total questions answered: %s \n", counter.Value(), questionCount.Value() )
	}()

	for i :=0;i<10 ;i++  {
		fmt.Println("What is the answer to: " + problems[i].Question())
		text, _ := reader2.ReadString('\n')
		if strings.TrimSpace(text) == "x" {
			fmt.Println("Quitting.")
		break
		} else if (strings.TrimSpace(text) == problems[i].Answer()) {
			fmt.Println("That's correct!")
			counter.Inc()
			questionCount.Inc()
		} else {
			questionCount.Inc()
			fmt.Println("Sorry, the correct answer is: " + problems[i].Answer())
		}
	}
	fmt.Printf("Correct answers: %s, total questions answered: %s \n", counter.Value(), questionCount.Value() )
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