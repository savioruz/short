package handler

import (
	"github.com/gofiber/fiber/v2/log"
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
		log.Fatalf("Error loading config: %v", err)
	}

	s := server.NewFiberServer(conf)
	s.Adaptor().ServeHTTP(w, r)
}
