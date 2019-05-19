package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var messageChannel = make(chan string, 100)

func channelConsumer() {
	message := <-messageChannel
	fmt.Println("request body:", message)
}

func handlePost(rw http.ResponseWriter, request *http.Request) {
	if body, err := ioutil.ReadAll(request.Body); err != nil {
		panic(err)
	} else {
		requestBody := string(body)
		messageChannel <- requestBody
	}
}

func main() {
	go channelConsumer()

	http.HandleFunc("/newtask", handlePost)
	http.ListenAndServe(":60730", nil)
}
