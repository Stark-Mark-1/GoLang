package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main(){
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'questions and answers'")
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err !=nil {
		exit("Failed to open the file")
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err !=nil {
		exit("Failed to open the file")
	}
	problems := parseLines(lines)
	correct := 0
	for i, p:= range problems{
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			fmt.Println("Correct!")
			correct++
		}
	}
}
	func parseLines(lines [][]string) []problem{
	ret := make([]problem, len(lines))
	for i, line := range lines{
		ret[i] = problem{
			q: line[0]
			a: line[1]
		}
	}
}

type problem struct{
	q string
	a string
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}
