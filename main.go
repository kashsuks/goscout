package main

import (
    "fmt"
    "flag"
	"net/http"
	"io"
)

func fetchTeamData(number int) string {
	url := fmt.Sprintf("https://api.ftcscout.org/rest/v1/teams/%v", number)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(resp.Body)
	data := string(body)
	return data
}

func main() {
	var teamNumber int
	flag.IntVar(&teamNumber, "number", 0, "team number flag")

	flag.Parse()
	
	//print out the team number => for future this will be checking if the team actually exists
	if teamNumber == 0 {
		fmt.Println("Pick a real team number")
	} else {
		fmt.Println(fetchTeamData(teamNumber))
	}
}
