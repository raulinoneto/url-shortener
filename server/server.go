// Package server serves HTTP Server
package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type (
	Options struct {
		Logger           logrus.FieldLogger
		MiddlewareSetter func(e *echo.Echo)
		HandlersSetter   func(e *echo.Echo)
		Port             string
		Host             string
	}

	Server struct {
		E               *echo.Echo
		port            string
		host            string
		serverRunning   bool
	}
)

func New(opt *Options) *Server {
	e := echo.New()
	e.Validator = NewCustomValidator()
	opt.MiddlewareSetter(e)
	opt.HandlersSetter(e)

	return &Server{
		E:             e,
		port:          opt.Port,
		host:          opt.Host,
		serverRunning: false,
	}
}

func (s *Server) quitSignal() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	var timeout time.Duration = 25

	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	s.E.Logger.Fatal("Waiting for the application to shutdown gracefully")

	if err := s.E.Shutdown(ctx); err != nil {
		s.E.Logger.Fatal(err)
	}

	s.E.Logger.Fatal("Application shutdown")
}

func (s *Server) serve() {
	if err := s.E.Start(s.host + ":" + s.port); err != nil {
		s.E.Logger.Fatal("Server could not start: " + err.Error())
	}

	s.E.Logger.Info("Server already started")
}

func (s *Server) Run() {
	go s.serve()
	s.serverRunning = true
	s.quitSignal()
}

func (s *Server) IsRunning() bool {
	return s.serverRunning
}

