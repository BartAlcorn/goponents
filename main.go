package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/bartalcorn/goponents/pkg/application"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := application.App.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
