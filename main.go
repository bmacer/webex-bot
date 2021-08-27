package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Url struct {
	Url string `json:"url"`
}

type PostMessage struct {
	RoomId string `json:"roomId"`
	Text   string `json:"text"`
}

type Messages struct {
	Items []Message `json:"items"`
}

type Message struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

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

type Extractable interface {
}

func extract(b []byte, ex Extractable) {
	err := json.Unmarshal(b, ex)
	if err != nil {
		fmt.Println("extract error: ", err)
	}
}

func sendRequest(reqType string, url string, b io.Reader) []byte {
	token := os.Getenv("WEBEX_API")
	client := &http.Client{}
	req, _ := http.NewRequest(reqType, url, b)
	t := fmt.Sprintf("Bearer %v", token)
	req.Header.Set("Authorization", t)
	req.Header.Set("Content-Type", "application/json")
	resp, e := client.Do(req)
	if e != nil {
		fmt.Println("client.Do error ", e)
	}
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		fmt.Println("ioutil.ReadAll error: ", e)
	}
	if resp.StatusCode > 300 {
		fmt.Println(url)
		fmt.Println("Status Code: ", resp.StatusCode)
		fmt.Println(string(body))
	}
	return body
}

func getRooms() Room {
	var r Room
	rooms_url := "https://webexapis.com/v1/rooms"
	r1 := sendRequest("GET", rooms_url, nil)
	extract(r1, &r)
	return r
}

func createRoom(roomName string) RoomItem {
	var r RoomItem
	rc := []byte(fmt.Sprintf("{\"title\": \"%v\"}", roomName))
	rooms_url := "https://webexapis.com/v1/rooms"
	r1 := sendRequest("POST", rooms_url, bytes.NewBuffer(rc))
	extract(r1, &r)
	return r
}

func deleteRoom(roomId string) []byte {
	rooms_url := fmt.Sprintf("https://webexapis.com/v1/rooms/%v", roomId)
	req := sendRequest("DELETE", rooms_url, nil)
	return req
}

func getMessagesFromRoom(roomId string) Messages {
	var m Messages
	url := fmt.Sprintf("https://webexapis.com/v1/messages?roomId=%v", roomId)
	req := sendRequest("GET", url, nil)
	extract(req, &m)
	return m
}

type People struct {
	Items []Person `json:"items"`
}

type Person struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
}

func queryPeopleByEmail(email string) People {
	var p People
	url := fmt.Sprintf("https://webexapis.com/v1/people?email=%v", email)
	res := sendRequest("GET", url, nil)
	extract(res, &p)
	return p
}

type Created struct {
	Created string `json:"created"`
}

func createMembershipById(roomId string, personId string) Created {
	var c Created
	mp := []byte(fmt.Sprintf("{\"roomId\": \"%v\", \"personId\": \"%v\"}",
		roomId, personId))
	membership_url := "https://webexapis.com/v1/memberships"
	res := sendRequest("POST", membership_url, bytes.NewBuffer(mp))
	extract(res, &c)
	return c
}

func postMessageToRoom(roomId string, msg string) Message {
	var m Message
	url := "https://webexapis.com/v1/messages"
	mp := []byte(fmt.Sprintf("{\"roomId\": \"%v\", \"text\": \"%v\"}",
		roomId, msg))
	req := sendRequest("POST", url, bytes.NewBuffer(mp))
	extract(req, &m)
	return m
}

func main() {
	/// My direct room
	direct := os.Getenv("WEBEX_TESTING_ROOM")

	/// Get bearer token API key
	var bearerToken string = os.Getenv("WEBEX_API")
	if bearerToken == "" {
		log.Fatal("Program requires ENV variable WEBEX_API with your bearer token")
	}

	/// Create a room, which returns a RoomItem that contains the id (c.Id)
	// c := createRoom("Bot room 1")

	/// Query by email, returns list of results,
	// u := queryPeopleByEmail("bmacer@cisco.com")
	// if len(u.Items) < 1 {
	// 	fmt.Println("No user exists")
	// 	os.Exit(1)
	// }
	// userToAdd := u.Items[0]

	/// Add a user (via user ID) to a room ("creating membership to a room")
	/// Returns Created object which has .Created timestamp
	// m := createMembershipById(c.Id, userToAdd.Id)
	// fmt.Println(m.Created)

	/// Get all rooms to delete all rooms matching a substring.  deleteRoom returns status code
	// for _, room := range getRooms().Items {
	// 	fmt.Printf("Room ID: %v | Room Title: %v\n", room.Id, room.Title)
	// 	if strings.Contains(room.Title, "bot and") {
	// 		deleteRoom(room.Id)
	// 	}
	// }

	/// Get all rooms and iterate through them to find a room
	// var room RoomItem
	// rooms := getRooms()
	// for _, r := range rooms.Items {
	// 	fmt.Printf("Room ID: %v | Room Title: %v\n", r.Id, r.Title)
	// 	if strings.Contains(r.Title, "Bot room 1") {
	// 		room = r
	// 	}
	// }

	/// Post a message to a room, returns MessageItem object
	/// which has Id and Text
	// m := postMessageToRoom(direct, "four!")
	// fmt.Println(m.Text)

	/// Get messages from a room
	msgs := getMessagesFromRoom(direct)
	for _, msg := range msgs.Items {
		fmt.Println(msg.Text)
	}
}
