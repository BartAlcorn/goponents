package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/bartalcorn/goponents/pkg/styles"
	"github.com/labstack/gommon/color"
)

type AppConfig struct {
	config Config
	Ctx    context.Context
	router http.Handler
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
	url := fmt.Sprintf("http://localhost:%v", a.config.ServerPort)
	fmt.Println(styles.StartBlockTop.Render(fmt.Sprintf("GOponents starting...")))
	fmt.Println(styles.StartBlockBottom.Render("Available at :", color.Yellow(color.Underline(url))))
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
