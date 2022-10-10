package main

import (
	"encoding/csv"
	"flag"
	"os"
	"log"
	"fmt"
)

func main(){

	// println("Hello world!")
	
	csvfilename := flag.String(
		"csvfile", // name
		"problems.csv", // value
		"The filename of the CSV file containing questions") // usage 	

	flag.Parse()

	file, err := os.Open(*csvfilename)
	if err != nil {
		log.Fatalf("\nError opening file:\n%s", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil{
		log.Fatalf("\nError reading CSV file:\n%s", err)
	}
	correct_answers := 0
	for _, value := range data {
		question := value[0]
		answer := value[1]
		//question , answer := value
		//fmt.Printf("Question: %s .... Answer: %s\n", question, answer)

		fmt.Printf("Question: %s\n", question)

		var input string
		fmt.Scanln(&input)

		//fmt.Printf("... you answered: %s\n", input)

		if input == answer{
			fmt.Println("Correct!")
			correct_answers ++
		} else {
			fmt.Println("Incorrect!")
		} 
	}

	fmt.Printf("All done! You have answered %d/%d questions correctly.",
		correct_answers, len(data))
	 //println(*csvfilename) // the * returns the value from the flag.String()
}