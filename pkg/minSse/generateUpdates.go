package minsse

import (
	"context"
	"math/rand"
	"time"

	"github.com/bartalcorn/excuses"
	"github.com/labstack/gommon/color"
)

func generateHistory(status string, a Asset) Asset {
	l := len(a.History)
	if l > 0 {
		a.History[0].End = time.Now()
	}
	var h History
	h.Process = tasks[rand.Intn(len(tasks))]
	h.Status = status
	h.Start = time.Now()
	if status == "error" {
		h.Description = excuses.Tech()
	}
	a.History = append([]History{h}, a.History...)
	a = generateElapsed(a)
	return a
}

func generateElapsed(a Asset) Asset {
	l := len(a.History)
	if l == 0 {
		return a
	}
	a.History[1].End = time.Now()
	a.History[1].Elapsed = a.History[1].End.Sub(a.History[1].Start).Round(3 * time.Millisecond)
	return a
}

func removeAsset(i int) {
	Assets = append(Assets[:i], Assets[i+1:]...)
}

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

	close(updateCh)

	// color.Println(color.Yellow("generateUpdates: Completed"))
}
