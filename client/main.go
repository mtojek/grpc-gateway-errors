package client

import (
	"bytes"
	"encoding/json"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"log"
	"net/http"
)

func CallServer() error {
	resp, err := http.Post("http://localhost:50001/v1/sayhello", "application/json",
		bytes.NewBuffer([]byte(`{"name": "Marcin"}`)))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusNotImplemented {
		log.Fatalf("Expected: http.StatusNotImplemented, Got: %v", resp.StatusCode)
	}

	var spbs spb.Status
	err = json.Unmarshal(body, &spbs)
	if err != nil {
		log.Fatal(err)
	}

	return status.ErrorProto(&spbs)
}