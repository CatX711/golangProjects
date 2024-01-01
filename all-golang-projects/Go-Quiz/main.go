package main

import "fmt"
import "flag"
import "os"

func main(){
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of `question,answer`")
	flag.Parse()

	// error check
	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("Failed to open the CSV file named: %s\n", *csvFilename)
		os.Exit(1)
	}

	_ = file
}	