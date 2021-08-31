package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bmacer/webex-bot/api"
)

func init() {
	fmt.Println("server package initialized")
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	var resp api.WebhookResponse
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &resp)
	if err != nil {
		fmt.Println("extract error: ", err)
	}

	y := api.GetMessage(resp.Data.Id)
	fmt.Println(y.Text)
	fmt.Println(y.PersonEmail)

}
