package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseToMap(dat string) (map[string]string, []string) {
	quizMap := make(map[string]string)
	var questions []string
	var in string
	var lastQ string
	for _, d := range dat {
		val := string(d)
		if val == "," {
			lastQ = in
			questions = append(questions, lastQ)
			in = ""
		} else if val == "\n" && lastQ != "" {
			quizMap[lastQ] = in
			lastQ = ""
			in = ""
		} else {
			in += val
		}
	}
	return quizMap, questions
}

func askQ(quizSol string, question string, numQ int, sec int) bool {
	timerDone := make(chan bool)
	var text string

	// We print the question and input the user response
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Question #%d: %s = \n", numQ, question)

	go func() {
		text, _ = reader.ReadString('\n')
		timerDone <- false
	}()

	go func() {
		time.Sleep(time.Duration(sec) * time.Second)
		timerDone <- true
	}()

	result := <-timerDone

	if result == true {
		return false
	} else {
		// We return if it is correct or not
		return (strings.TrimSpace(quizSol) == strings.TrimSpace(text))
	}
}

func main() {
	// We check for a valid .csv file format
	csvFilename := flag.String("csv", "problems.csv", "valid format is [question], [answer]")
	// We check for a valid seconds arg format
	secArg := flag.Int("timer", 2, "seconds to answer each question")

	flag.Parse()
	_ = csvFilename
	_ = secArg

	// We read the .csv file
	dat, err := ioutil.ReadFile(os.Args[1])
	check(err)

	// We take the seconds input from the user
	sec, err := strconv.Atoi(os.Args[2])
	check(err)

	// We parse the problems
	quizMap, questions := parseToMap(string(dat))

	// We set the variables for the final score
	total := len(questions)
	score := 0

	// For each question
	for i := 0; i < len(questions); i++ {
		// We ask each question to the user
		res := askQ(quizMap[questions[i]], questions[i], i+1, sec)
		// We record the question results (correct/false)
		if res == true {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Println("Incorrect")
		}
	}

	// We print out the final score
	fmt.Printf("%d/%d\n", score, total)
}
