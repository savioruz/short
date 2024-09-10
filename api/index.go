package handler

import (
	"github.com/savioruz/short/config"
	_ "github.com/savioruz/short/docs"
	"github.com/savioruz/short/pkg/server"
	"net/http"
)

// Handler is a function main for vercel serverless
func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()

	conf, err := config.LoadConfig()
	if err != nil {
		http.Error(w, "Error loading config", http.StatusInternalServerError)
		return
	}

	s := server.NewFiberServer(conf)
	s.Adaptor().ServeHTTP(w, r)
}
