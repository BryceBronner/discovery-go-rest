package main

import (
	"encoding/json"
	"net/http"
)

type RequestContext struct {
	writer  http.ResponseWriter
	request *http.Request
}

func (context *RequestContext) JSON(status int, payload interface{}) {
	context.writer.WriteHeader(status)
	context.writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(context.writer).Encode(payload)
}
