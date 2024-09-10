package server

import "net/http"

type Server interface {
	ServerStart()
	Adaptor() http.HandlerFunc
}
