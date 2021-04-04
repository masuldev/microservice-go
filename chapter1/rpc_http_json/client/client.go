package client

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/masuldev/micro-go/chapter1/rpc_http_json/contract"
)

func PerformRequest() contract.HelloWorldResponse {
	r, _ := http.Post(
		"http://localhost:1234",
			"application/json",
			bytes.NewBuffer([]byte(`{"id": 1, "method": "HelloWorldHandler.HelloWorld", "params": [{"name":"World"}]}`)),
		)
	defer r.Body.Close()

	var response contract.HelloWorldResponse
	json.NewDecoder(r.Body).Decode(&response)

	return response
}
