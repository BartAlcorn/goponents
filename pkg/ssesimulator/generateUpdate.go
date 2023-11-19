package ssesimulator

import (
	"context"
	"time"

	"github.com/labstack/gommon/color"
)

func generateUpdate(ctx context.Context, minEventCh chan<- Asset) {
	ticker := time.NewTicker(time.Second)
	count := 0

dataLoop:
	for {
		select {
		case <-ctx.Done():
			break dataLoop
		case <-ticker.C:
			minEventCh <- MakeAsset()
			count += 1
			if count > 3 {
				break dataLoop
			}
		}
	}

	ticker.Stop()

	close(minEventCh)

	color.Println(color.Yellow("generateUpdate: Completed"))

}
