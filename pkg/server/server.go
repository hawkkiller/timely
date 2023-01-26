package server

import (
	"errors"
	"io"
	"net/http"
	"os"
	"syscall"

	"example.com/template/pkg/logger"
	"example.com/template/pkg/shutdown"
)

func Start(handler http.Handler, log logger.Logger, closables ...io.Closer) {
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	signals := []os.Signal{syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM}
	closables = append(closables, server)

	go shutdown.Graceful(signals, closables...)
	err := server.ListenAndServe()

	if err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			log.Info("Server shutdown")
		default:
			log.Fatalf("Server error: %v", err)
		}
	}
}
