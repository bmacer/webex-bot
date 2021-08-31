package api

import (
	"bytes"
	"encoding/json"
)

type PostMessage struct {
	RoomId      string                   `json:"roomId"`
	Text        string                   `json:"text"`
	Attachments []AdaptiveCardAttachment `json:"attachments"`
}

type Messages struct {
	Items []Message `json:"items"`
}

type Message struct {
	Id          string                   `json:"id"`
	Text        string                   `json:"text"`
	PersonId    string                   `json:"personId"`
	PersonEmail string                   `json:"personEmail"`
	Attachments []AdaptiveCardAttachment `json:"attachments"`
}

type AdaptiveCardAttachment struct {
	ContentType string              `json:"contentType"`
	Content     AdaptiveCardContent `json:"content"`
}

type AdaptiveCardContent struct {
	Type    string                     `json:"type"`
	Body    []AdaptiveContentTextBlock `json:"body"`
	Schema  string                     `json:"$schema"`
	Version string                     `json:"version"`
}

type AdaptiveContentTextBlock struct {
	Type string `json:"type"`
	Text string `json:"text"`
	Wrap bool   `json:"wrap"`
}

func NewSimpleAttachment(msg string) []AdaptiveCardAttachment {
	content := AdaptiveCardContent{
		Type: "AdaptiveCard",
		Body: []AdaptiveContentTextBlock{
			AdaptiveContentTextBlock{
				Type: "TextBlock",
				Text: msg,
				Wrap: true,
			},
		},
		Schema:  "http://adaptivecards.io/schemas/adaptive-card.json",
		Version: "1.2",
	}

	attachment := []AdaptiveCardAttachment{
		AdaptiveCardAttachment{
			ContentType: "application/vnd.microsoft.card.adaptive",
			Content:     content,
		},
	}

	return attachment
}

func NewMessage(roomId string, msg string) PostMessage {
	return PostMessage{
		RoomId: roomId,
		Text:   msg,
	}

}

func NewMessageWithAttachment(roomId string, msg string, testMessage string) PostMessage {
	return PostMessage{
		RoomId:      roomId,
		Text:        msg,
		Attachments: NewSimpleAttachment(testMessage),
	}
}

func PostMessageWithAdaptiveCard(p PostMessage) Message {
	var m Message
	url := "https://webexapis.com/v1/messages"
	j, _ := json.Marshal(p)
	mp := []byte(j)
	req := sendRequest("POST", url, bytes.NewBuffer(mp))
	extract(req, &m)
	return m

}

// func PostMessageToRoom(roomId string, msg string) Message {
// 	var m Message
// 	url := "https://webexapis.com/v1/messages"

// 	j, _ := json.Marshal(GetSimpleAttachment())

// 	c := fmt.Sprintf(`
// 	{
// 		"roomId": "%v",
// 		"text": "%v",
// 		"attachments": %v
// 	}`, roomId, msg, string(j))

// 	fmt.Println(c)

// 	mp := []byte(c)
// 	req := sendRequest("POST", url, bytes.NewBuffer(mp))
// 	extract(req, &m)
// 	return m
// }
