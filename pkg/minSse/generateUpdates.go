package minsse

import (
	"context"
	"time"

	"github.com/labstack/gommon/color"
)

func removeAsset(i int) {
	Assets = append(Assets[:i], Assets[i+1:]...)
}

// generateAssets responds to a ticker, periodically
// creating an Update to an existing Asset
// and pushes the update onto the updateCh channel
func generateUpdates(ctx context.Context, updateCh chan<- Asset) {
	ticker := time.NewTicker(500 * time.Millisecond)
	// color.Println(color.Green("generateUpdates: started"))

dataLoop:
	for {
		select {
		case <-ctx.Done():
			break dataLoop
		case <-ticker.C:
			if len(Assets) < 1 {
				break dataLoop
			}
			l := len(Assets) - 1
			if l < 0 {
				color.Print(color.Red("index out of bounds"))
				break dataLoop
			}
			go MakeUpdate(ctx, updateCh)
		}
	}

	ticker.Stop()

	// close(updateCh)

	// color.Println(color.Yellow("generateUpdates: Completed"))
}
