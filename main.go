package main

import (
	"log"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	router := createRouter()
	router.GET("/", func(context *RequestContext) {
		payload := Message{Message: "This example work with request context"}
		context.JSON(http.StatusOK, payload)
	})
	log.Fatal(http.ListenAndServe(":8080", router))
}
