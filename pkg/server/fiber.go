package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/savioruz/short/config"
	"github.com/savioruz/short/internal/adapters/cache"
	"github.com/savioruz/short/internal/adapters/handlers/rest"
	"github.com/savioruz/short/internal/adapters/repositories"
	"github.com/savioruz/short/internal/cores/services"
	"github.com/savioruz/short/pkg/middlewares"
	"github.com/savioruz/short/pkg/routes"
	"net/http"
	"os"
	"os/signal"
)

type Fiber struct {
	app  *fiber.App
	conf *config.Config
}

// NewFiberServer is a function to initialize the fiber server
func NewFiberServer(conf *config.Config) Fiber {
	a := fiber.New()

	// Middleware
	middlewares.FiberMiddleware(a)
	middlewares.LimiterMiddleware(a)
	middlewares.MonitorMiddleware(a)

	return Fiber{
		app:  a,
		conf: conf,
	}
}

func (s *Fiber) ServerStart() {
	s.initializeShortURLHandler()
	s.initializeRoutes()
	s.startServerWithGrafeculShutdown()
}

func (s *Fiber) Adaptor() http.HandlerFunc {
	s.initializeShortURLHandler()
	s.initializeRoutes()

	return adaptor.FiberApp(s.app)
}

func (s *Fiber) initializeRoutes() {
	routes.SwaggerRoute(s.app)
	routes.NotFoundRoute(s.app)
}

func (s *Fiber) initializeShortURLHandler() {
	redis, err := cache.NewRedisCache(s.conf.Redis.Addr, s.conf.Redis.Password, s.conf.Redis.DB)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	repos := repositories.NewDB(redis)
	service := services.NewShortURLService(repos)
	handler := rest.NewShortURLHandler(service)

	r := s.app.Group("/api/v1")
	r.Post("/shorten", handler.CreateShortURL)

	s.app.Get("/:url", handler.ResolveURL)
}

func (s *Fiber) startServerWithGrafeculShutdown() {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := s.app.Shutdown(); err != nil {
			log.Errorf("Fiber shutdown error: %v", err)
		}

		close(idleConnsClosed)
	}()

	fiberConnectionURL := fmt.Sprintf("%s:%s", s.conf.Server.Host, s.conf.Server.Port)

	if err := s.app.Listen(fiberConnectionURL); err != nil {
		log.Errorf("Fiber listen error: %v", err)
	}

	<-idleConnsClosed
}
