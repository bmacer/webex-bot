package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	api "github.com/bmacer/webex-bot/api"
	"github.com/bmacer/webex-bot/server"
)

func main() {

	ac, _ := ioutil.ReadFile("card copy.json")
	acs := string(ac)
	acs = strings.Replace(acs, "REPLACEME", "GHI", 1)
	fmt.Println("acs:", acs)

	aca := api.AdaptiveCardAttachment{
		ContentType: "application/vnd.microsoft.card.adaptive",
		Content:     acs,
	}

	api.PostMessageWithAdaptiveCard(api.GetRooms().Items[1].Id, aca)
	os.Exit(0)

	http.HandleFunc("/", server.RequestHandler)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err)
	}
}
