package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type AppConfig struct {
	router  http.Handler
	RedisDB *redis.Client
	config  Config
}

var App *AppConfig

func init() {
	config := LoadConfig()
	App = &AppConfig{
		RedisDB: redis.NewClient(&redis.Options{
			Addr: config.RedisAddress,
		}),
		config: config,
	}

	App.loadRoutes()
}

func (a *AppConfig) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.ServerPort),
		Handler: a.router,
	}
	fmt.Println("Starting on PORT", a.config.ServerPort)

	err := a.RedisDB.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	defer func() {
		if err := a.RedisDB.Close(); err != nil {
			fmt.Println("failed to close redis", err)
		}
	}()

	fmt.Println("Starting server")

	ch := make(chan error, 1)

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}

	// return nil
}