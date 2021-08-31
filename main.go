package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	api "github.com/bmacer/webex-bot/api"
)

func main() {
	msg := api.NewMessageWithAttachment(
		api.GetRooms().Items[1].Id,
		"main message",
		"inside messagee",
	)
	x := api.PostMessageWithAdaptiveCard(msg)
	fmt.Println(x)
	// api.PostMessageToRoom(api.GetRooms().Items[1].Id, "abcde")

	// created_webhook := api.CreateWebhook(
	// 	"mywebhookIP", "http://anyaelyse.comm:8888", "messages", api.WebhookEventCreated,
	// )
	// fmt.Println(created_webhook)

	// wh := api.GetWebhooks()
	// for _, w := range wh.Items {
	// 	fmt.Println(w)
	// }
	os.Exit(0)

	http.HandleFunc("/", RequestHandler)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type WebhookResponseData struct {
	Id          string `json:"id"`
	RoomId      string `json:"roomId"`
	RoomType    string `json:"roomType"`
	PersonId    string `json:"personId"`
	PersonEmail string `json:"personEmail"`
	Created     string `json:"created"`
}

type WebhookResponse struct {
	Id        string              `json:"id"`
	Name      string              `json:"name"`
	TargetUrl string              `json:"targetUrl"`
	Resource  string              `json:"resource"`
	Event     string              `json:"event"`
	OrgId     string              `json:"ordId"`
	CreatedBy string              `json:"createdBy"`
	AppId     string              `json:"appId"`
	OwnedBy   string              `json:"ownedBy"`
	Status    string              `json:"status"`
	Created   string              `json:"created"`
	ActorId   string              `json:"actorId"`
	Data      WebhookResponseData `json:"data"`
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	var resp WebhookResponse
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &resp)
	if err != nil {
		fmt.Println("extract error: ", err)
	}

	y := api.GetMessage(resp.Data.Id)
	fmt.Println(y.Text)
	fmt.Println(y.PersonEmail)

}
