package ssesimulator

import (
	"context"
	"fmt"
	"time"
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

	fmt.Println("generateUpdate: Completed")
}
