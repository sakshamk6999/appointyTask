package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateContactError(t *testing.T) {
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

func TestNullListContact(t *testing.T) {
	req, err := http.NewRequest("GET", "/contacts?user=10&infection_timestamp=1607731200", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListContactsHandler)

	handler.ServeHTTP(responseRecorder, req)
	fmt.Println(responseRecorder.Body)
	status := responseRecorder.Code
	if status != 200 {
		t.Errorf("status code does not match !, got %v", status)
	}

}
