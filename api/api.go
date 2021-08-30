package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func init() {
	fmt.Println("api package initialized")
}

type Url struct {
	Url string `json:"url"`
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

func GetMessage(messageId string) Message {
	var m Message
	url := fmt.Sprintf("https://webexapis.com/v1/messages/%v", messageId)
	fmt.Println(url) //TODO delete
	res := sendRequest("GET", url, nil)
	extract(res, &m)
	return m
}

func GetMessagesFromRoom(roomId string) Messages {
	var m Messages
	url := fmt.Sprintf("https://webexapis.com/v1/messages?roomId=%v", roomId)
	req := sendRequest("GET", url, nil)
	extract(req, &m)
	return m
}

/*
func main() {
	/// My direct room
	direct := os.Getenv("WEBEX_TESTING_ROOM")

	/// Get bearer token API key
	var bearerToken string = os.Getenv("WEBEX_API")
	if bearerToken == "" {
		log.Fatal("Program requires ENV variable WEBEX_API with your bearer token")
	}

	/// Get all webhooks
	// for _, wh := range GetWebhooks().Items {
	// 	fmt.Println(wh)
	// }

	/// Create a webhook
	// created_webhook := CreateWebhook(
	// 	"mywebhookIP", "http://13.82.176.69:8888", "messages", WebhookEventCreated,
	// )
	// fmt.Println(created_webhook)

	/// Delete a webhook
	// webhookId := "Y2lzY29zcGFyazovL3VzL1dFQkhPT0svMTJhMDE5MGMtYTU0Yy00ZjFjLTlkYjctOTM5YjViMjBkMTUy"
	// res := DeleteWebhook(webhookId)
	// fmt.Println(string(res)) // Blank == OK, error == not OK

	/// Delete all created webhooks
	// for _, wh := range GetWebhooks().Items {
	// 	DeleteWebhook(wh.Id)
	// }

	/// Create a room, which returns a RoomItem that contains the id (c.Id)
	// c := CreateRoom("Bot room 1")

	/// Query by email, returns list of results,
	// u := QueryPeopleByEmail("bmacer@cisco.com")
	// if len(u.Items) < 1 {
	// 	fmt.Println("No user exists")
	// 	os.Exit(1)
	// }
	// userToAdd := u.Items[0]

	/// Add a user (via user ID) to a room ("creating membership to a room")
	/// Returns Created object which has .Created timestamp
	// m := CreateMembershipById(c.Id, userToAdd.Id)
	// fmt.Println(m.Created)

	/// Get all rooms to delete all rooms matching a substring.  deleteRoom returns status code
	// for _, room := range GetRooms().Items {
	// 	fmt.Printf("Room ID: %v | Room Title: %v\n", room.Id, room.Title)
	// 	if strings.Contains(room.Title, "bot and") {
	// 		DeleteRoom(room.Id)
	// 	}
	// }

	/// Get all rooms and iterate through them to find a room
	// var room RoomItem
	// rooms := GetRooms()
	// for _, r := range rooms.Items {
	// 	fmt.Printf("Room ID: %v | Room Title: %v\n", r.Id, r.Title)
	// 	if strings.Contains(r.Title, "Bot room 1") {
	// 		room = r
	// 	}
	// }

	/// Post a message to a room, returns MessageItem object
	/// which has Id and Text
	// m := PostMessageToRoom(direct, "four!")
	// fmt.Println(m.Text)

	/// Get messages from a room
	msgs := GetMessagesFromRoom(direct)
	for _, msg := range msgs.Items {
		fmt.Println(msg.Text)
	}
	var _ = direct
}
*/
