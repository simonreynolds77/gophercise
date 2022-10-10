package main

import "flag"

func main(){
	
	csvfilename := flag.String(
		"csvfile", 
		"problems.csv", 
		"The filename of the CSV file containing questions")
	
	flag.Parse()

	println(*csvfilename)
}