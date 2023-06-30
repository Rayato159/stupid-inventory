package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Rayato159/stupid-inventory/src/config"
	"github.com/Rayato159/stupid-inventory/src/controllers/httpControllers"
	"github.com/Rayato159/stupid-inventory/src/repositories"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type IHttpServer interface {
	StartUserServer()
	StartItemServer()
}

type httpServer struct {
	app      *echo.Echo
	cfg      *config.Config
	dbClient *mongo.Client
}

func NewHttpServer(cfg *config.Config, dbClient *mongo.Client) IHttpServer {
	return &httpServer{
		app:      echo.New(),
		cfg:      cfg,
		dbClient: dbClient,
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

	log.Println("server is shutting down")
	if err := s.app.Shutdown(ctx); err != nil {
		s.app.Logger.Fatal(err)
	}
}

func (s *httpServer) StartUserServer() {
	userController := &httpControllers.UserHttpController{
		UserRepository: &repositories.UserRepository{
			Client: s.dbClient,
		},
	}

	s.app.GET("/user/:user_id", userController.FindOneUser)

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
