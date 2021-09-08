package api

import (
	"bytes"
	"fmt"
)

type MessagePosting struct {
	RoomId      string                   `json:"roomId"`
	Text        string                   `json:"text"`
	Attachments []AdaptiveCardAttachment `json:"attachments"`
}

type Messages struct {
	Items []Message `json:"items"`
}

type Message struct {
	Id          string `json:"id"`
	Text        string `json:"text"`
	PersonId    string `json:"personId"`
	PersonEmail string `json:"personEmail"`
	Attachments string `json:"attachments"`
}

func (aca *AdaptiveCardAttachment) AsBytes() []byte {
	var b []byte
	b = append(b, []byte(fmt.Sprintf(`{"contentType":"%v","content":`, aca.ContentType))...)
	b = append(b, aca.Content...)
	b = append(b, byte('}'))
	return b
}

type AdaptiveCardAttachment struct {
	ContentType string `json:"contentType"`
	Content     string `json:"content"`
}

func PostMessage(roomId string, msg string) Message {
	var m Message
	url := "https://webexapis.com/v1/messages"

	c := fmt.Sprintf(`{"roomId": "%v","text": "%v"}`, roomId, msg)
	fmt.Println(c)
	mp := []byte(c)
	req := sendRequest("POST", url, bytes.NewBuffer(mp))
	extract(req, &m)
	return m
}

func PostMessageWithAdaptiveCard(roomId string, aca AdaptiveCardAttachment) Message {
	var m Message
	url := "https://webexapis.com/v1/messages"
	c := fmt.Sprintf(`
	{
		"roomId": "%v",
		"text": "nothing",
		"attachments": %v
	}`, roomId, string(aca.AsBytes()))
	mp := []byte(c)
	req := sendRequest("POST", url, bytes.NewBuffer(mp))
	fmt.Printf("type: %T\n", m)
	fmt.Println(string(req))
	extract(req, &m)
	return m
}
