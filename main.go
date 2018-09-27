package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main(){

	csvFileName := flag.String("csv", "problems.csv",
		"A csv file in the format of a question/answer format")

	timeLimit := flag.Int("Limit", 30, "the time limit for the quiz in seconds")

	flag.Parse()

	_ = csvFileName

	file, err := os.Open(*csvFileName)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file %s\n", *csvFileName))
	}

	read := csv.NewReader(file)

	lines, err := read.ReadAll()

	if err != nil {
		exit("Failed to parse the provided CSV file")
	}

	Problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit)* time.Second)


	correct := 0

problemloop:
	for i, p := range Problems {
		fmt.Printf("Problem #%d: %s = ", i + 1, p.Question )

		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer

		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerCh:
			if answer == p.Answer {
				correct ++
			}

		}

	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(Problems))
}


func parseLines(lines [][]string) []Problems {

	ret := make([]Problems, len(lines))

	for i, line := range lines {
		ret[i] = Problems{
			Question: line[0],
			Answer: strings.TrimSpace(strings.ToLower(line[1])),

		}
	}

	return ret

}

type Problems struct {
	Question string
	Answer string
}

func exit(msg string){

	fmt.Println(msg)
	os.Exit(1)
}