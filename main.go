package main

import (
	"context"
	"github.com/raulinoneto/url-shortener/logger"
	"github.com/raulinoneto/url-shortener/server"
)

func main() {

	_, log := logger.Init(context.Background())

	opt := &server.Options{
		Logger: log,
		MiddlewareSetter: nil,
		HandlersSetter:   nil,
	} 
	s := server.New(opt)

	s.Run()
}
