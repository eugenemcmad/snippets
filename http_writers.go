package tests

import (
	"encoding/json"
	"net/http"
)

// Writes error as response in JSON format
func writeErr(wr http.ResponseWriter, err error) {
	bytes, _ := json.Marshal(struct {
		Error string
	}{err.Error()})
	wr.Write(bytes)
}

// Writes any struct as response in JSON format
func writeResp(wr http.ResponseWriter, resp interface{}) {
	bytes, err := json.Marshal(resp)
	if err != nil {
		writeErr(wr, err)
		return
	}
	wr.Write(bytes)
}
