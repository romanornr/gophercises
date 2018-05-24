package main

import (
	"flag"
	"os"
	"fmt"
	"encoding/csv"
	"time"
)

func main(){
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of question, answer")
	TimeLimit := flag.Int("limit", 30, "The time limit for this quiz in seconds.")
	flag.Parse()

	file, err := os.Open(*csvFileName) // csvFile name is a pointer to a string
	if err != nil{
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil{
		exit("failed to parse the provided CSV file")
	}
	fmt.Println(lines)

	problems := parseLines(lines)


	timer := time.NewTicker(time.Duration(*TimeLimit) * time.Second)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1,  p.question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("You scored %d out out %d.\n", correct, len(problems))
			return
		case answer := <-answerCh:
		if answer == p.answer {
			correct++
		}

		//default:
		//	fmt.Printf("Problem #%d: %s = \n", i+1,  p.question)
		//	var answer string
		//	fmt.Scanf("%s\n", &answer) // get rid of trailing spaces
		//	if answer == p.answer {
		//		fmt.Println("Correct!")
		//		correct++
		//	}
		}
	}
	fmt.Printf("You scored %d out out %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines)) //every row is a problem
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer: line[1],
		}
	}
	return ret
}

type problem struct {
	question string
	answer string
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}