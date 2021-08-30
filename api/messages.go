package api

import (
	"bytes"
	"fmt"
)

type PostMessage struct {
	RoomId string `json:"roomId"`
	Text   string `json:"text"`
}

type Messages struct {
	Items []Message `json:"items"`
}

type Message struct {
	Id          string `json:"id"`
	Text        string `json:"text"`
	PersonId    string `json:"personId"`
	PersonEmail string `json:"personEmail"`
}

func PostMessageToRoom(roomId string, msg string) Message {
	var m Message
	url := "https://webexapis.com/v1/messages"
	mp := []byte(fmt.Sprintf("{\"roomId\": \"%v\", \"text\": \"%v\"}",
		roomId, msg))
	req := sendRequest("POST", url, bytes.NewBuffer(mp))
	extract(req, &m)
	return m
}
