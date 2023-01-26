package main

import (
	"example.com/template/internal/router"
	"example.com/template/pkg/logger"
	"example.com/template/pkg/server"
)

// main is the entry point for the application.
func main() {
	logger.Init()
	log := logger.GetLogger()
	r := router.ConfigureRouter()
	server.Start(r, log)
}
