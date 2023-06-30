package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Rayato159/stupid-inventory/src/config"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type IHttpServer interface {
	StartUserServer()
	StartItemServer()
}

type httpServer struct {
	app *echo.Echo
	cfg *config.Config
	db  *mongo.Client
}

func NewHttpServer(cfg *config.Config, db *mongo.Client) IHttpServer {
	return &httpServer{
		app: echo.New(),
		cfg: cfg,
		db:  db,
	}
}

func (s *httpServer) listen() {
	log.Printf("%s server is starting on %s", s.cfg.App.AppName, s.cfg.App.Url)
	if err := s.app.Start(s.cfg.App.Url); err != nil && err != http.ErrServerClosed {
		s.app.Logger.Fatal("shutting down the server")
	}
}

// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
// Use a buffered channel to avoid missing signals as recommended for signal.Notify
func (s *httpServer) gracefullyShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.app.Shutdown(ctx); err != nil {
		s.app.Logger.Fatal(err)
	}
}

func (s *httpServer) StartUserServer() {
	s.app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Start server
	go s.listen()
	s.gracefullyShutdown()
}

func (s *httpServer) StartItemServer() {
	s.app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Start server
	go s.listen()
	s.gracefullyShutdown()
}
