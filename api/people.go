package api

import "fmt"

type People struct {
	Items []Person `json:"items"`
}

type Person struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
}

func QueryPeopleByEmail(email string) People {
	var p People
	url := fmt.Sprintf("https://webexapis.com/v1/people?email=%v", email)
	res := sendRequest("GET", url, nil)
	extract(res, &p)
	return p
}
