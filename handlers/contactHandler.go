package handlers

import "net/http"

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ListContactsHandler(w, r)
	} else {
		CreateContactHandler(w, r)
	}
}
