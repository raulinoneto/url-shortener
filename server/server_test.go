package server

import (
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestServer_Run(t *testing.T) {
	s := New(&Options{
		Port:             "4321",
		MiddlewareSetter: func(e *echo.Echo) {},
		HandlersSetter:   func(e *echo.Echo) {},
	})
	go s.Run()
	time.Sleep(2 * time.Second)
	assert.True(t, s.IsRunning(), "Server is not running")
}
