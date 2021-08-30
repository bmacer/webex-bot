package api

import (
	"bytes"
	"fmt"
)

type Created struct {
	Created string `json:"created"`
}

func CreateMembershipById(roomId string, personId string) Created {
	var c Created
	mp := []byte(fmt.Sprintf("{\"roomId\": \"%v\", \"personId\": \"%v\"}",
		roomId, personId))
	membership_url := "https://webexapis.com/v1/memberships"
	res := sendRequest("POST", membership_url, bytes.NewBuffer(mp))
	extract(res, &c)
	return c
}
