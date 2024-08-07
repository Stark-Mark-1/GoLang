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
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}
