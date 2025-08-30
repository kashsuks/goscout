package main

import (
    "fmt"
    "flag"
	// "net/http"
)

func main() {

	teamNumber := flag.Int("number", 0, "string flag")

	flag.Parse()
	
	if *teamNumber == 0 {
		fmt.Println("Pick a real team number")
	} else {
		fmt.Println("Team Number: ", *teamNumber)
	}
}