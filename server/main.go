package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Set routing rules
	http.HandleFunc("/main", Tmp0)
	http.HandleFunc("/", Tmp1)
	http.HandleFunc("/hello", Tmp2)

	//Use the default DefaultServeMux.
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Tmp0(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "version 0")
	fmt.Println("hello0")
}
func Tmp1(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "version 1")
	fmt.Println("hello1")
}

func Tmp2(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "version 2")
	fmt.Println("hello2")
}
