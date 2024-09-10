package handler

import (
	_ "github.com/savioruz/short/docs"
	"github.com/savioruz/short/pkg/server"
	"net/http"
)

// Handler is a function main for vercel serverless
func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()

	s := server.NewFiberAdaptor()
	s.Adaptor().ServeHTTP(w, r)
}
