package main

import (
    "fmt"
    "flag"
	// "net/http"
)

func main() {

	teamNumber := flag.Int("number", 0, "string flag")

	flag.Parse()
	
	//print out the team number => for future this will be checking if the team actually exists
	if *teamNumber == 0 {
		fmt.Println("Pick a real team number")
	} else {
		fmt.Println("Team Number: ", *teamNumber)
	}
}