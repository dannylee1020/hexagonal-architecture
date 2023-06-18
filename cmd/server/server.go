package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dannylee/url-ports-adapters/adapters"
)

func serve(config adapters.Config, handler http.Handler) error {
	srv := &http.Server{
		Addr:        fmt.Sprintf(":%d", config.Port),
		Handler:     handler,
		IdleTimeout: time.Minute,
	}

	errorChan := make(chan error)

	// run separate goroutine to listen to server shutdown
	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit

		log.Printf("Shutting down server with signal: %s", s.String())

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		errorChan <- srv.Shutdown(ctx)
	}()

	log.Printf("Starting server at: %v", config.Port)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-errorChan
	if err != nil {
		return err
	}

	log.Printf("Stopped server at :%v", config.Port)

	return nil
}
