package main

import (
	"io/ioutil"
	"strings"

	"github.com/bmacer/webex-bot/api"
)

func main() {
	// test_room := "Y2lzY29zcGFyazovL3VzL1JPT00vZjc0NzM0NDAtMGJmMy0xMWVjLWJmZjctOTNiMGQxYTk5MzEx"

	// api.PostMessage(test_room, "test")
	// for i := 0; i < 5; i++ {
	ac, _ := ioutil.ReadFile("card copy.json")
	acs := string(ac)
	// r := food.GetRandomRecipes()
	// text := fmt.Sprint(r.Recipes[i].Id) + " " + r.Recipes[i].Title
	text := "Replaced"
	acs = strings.Replace(acs, "1REPLACEME1", text, 1)
	acs = strings.Replace(acs, "2REPLACEME2", text, 1)
	aca := api.AdaptiveCardAttachment{
		ContentType: "application/vnd.microsoft.card.adaptive",
		Content:     acs,
	}
	api.PostMessageWithAdaptiveCard(api.GetRooms().Items[1].Id, aca)
	// }
	//
	// server.Run()

}
