package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	api "github.com/bmacer/webex-bot/api"
)

func main() {
	// Set routing rules
	http.HandleFunc("/", RequestHandler)
	// http.HandleFunc("/", Tmp1)
	// http.HandleFunc("/hello", Tmp2)
	//Use the default DefaultServeMux.
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
	// x := api.GetWebhooks()
	// fmt.Println(x.Items[0].Id)

	y := api.GetMessage(resp.Data.Id)
	fmt.Println(y.Text)
	fmt.Println(y.PersonEmail)

}
