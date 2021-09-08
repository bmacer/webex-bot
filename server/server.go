package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
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

func AdminHandler(w http.ResponseWriter, r *http.Request) {

	b, _ := ioutil.ReadAll(r.Body)
	io.WriteString(w, "print to page")
	fmt.Println(string(b))

}

func Run() {
	fs := http.FileServer(http.Dir("admin"))

	http.HandleFunc("/", RequestHandler)
	http.HandleFunc("/admin", AdminHandler)
	http.Handle("/static", fs)

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
