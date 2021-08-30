package api

import (
	"bytes"
	"fmt"
)

type Room struct {
	Items []RoomItem `json:"items"`
}

type RoomItem struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	Type         string `json:"type"`
	IsLocked     bool   `json:"isLocked"`
	LastActivity string `json:"lastActivity"`
	CreatorId    string `json:"creatorId"`
	Created      string `json:"created"`
	// OwnerId      string `json:"ownerId"`
}

type RoomCreation struct {
	Title string `json:"title"`
	// TeamId           string `json:"teamId"`
	// ClassificationId string `json:"classificationId"`
}

func GetRooms() Room {
	var r Room
	rooms_url := "https://webexapis.com/v1/rooms"
	r1 := sendRequest("GET", rooms_url, nil)
	extract(r1, &r)
	return r
}

func CreateRoom(roomName string) RoomItem {
	var r RoomItem
	rc := []byte(fmt.Sprintf("{\"title\": \"%v\"}", roomName))
	rooms_url := "https://webexapis.com/v1/rooms"
	r1 := sendRequest("POST", rooms_url, bytes.NewBuffer(rc))
	extract(r1, &r)
	return r
}

func DeleteRoom(roomId string) []byte {
	rooms_url := fmt.Sprintf("https://webexapis.com/v1/rooms/%v", roomId)
	req := sendRequest("DELETE", rooms_url, nil)
	return req
}
