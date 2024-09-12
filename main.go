package main

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/savioruz/short/config"
	_ "github.com/savioruz/short/docs"
	"github.com/savioruz/short/pkg/server"
)

// @title Short URL API
// @version 0.1
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email jakueenak@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	s := server.NewFiberServer(conf)
	s.ServerStart()
}
