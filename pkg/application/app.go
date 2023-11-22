package application

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type AppConfig struct {
	router http.Handler
	Ctx    context.Context
	config Config
}

var App *AppConfig

func init() {
	config := LoadConfig()
	App = &AppConfig{
		config: config,
	}

	App.loadRoutes()
}

func (a *AppConfig) Start(ctx context.Context) error {
	a.Ctx = ctx
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.ServerPort),
		Handler: a.router,
	}
	fmt.Println("Starting on PORT", a.config.ServerPort)
	fmt.Println("Starting server")

	ch := make(chan error, 1)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}

	// return nil
}
