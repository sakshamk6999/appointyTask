package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateContact(t *testing.T) {
	var jsonStr = []byte(`
	{
		"userIdOne":"10",
		"userIdTwo":"1",
		"contactTime": "2020-12-12T00:00:00+00:00"
	}
	`)
	fmt.Println("I am here")
	req, err := http.NewRequest("POST", "/contacts", bytes.NewBuffer(jsonStr))
	fmt.Println("in here")
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateContactHandler)

	handler.ServeHTTP(responseRecorder, req)
	fmt.Println(responseRecorder.Body.String())
}
