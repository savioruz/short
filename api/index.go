package handler

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()

}

func initializeHandler() {
	// Code here
}
