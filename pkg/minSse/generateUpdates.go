package minSse

import (
	"context"
	"math/rand"
	"time"

	"github.com/bartalcorn/excuses"
	"github.com/labstack/gommon/color"
)

func generateHistory(status string, a Asset) Asset {
	var h History
	h.Process = tasks[rand.Intn(len(tasks))]
	h.Status = status
	h.Start = time.Now()
	if status == "error" {
		h.Description = excuses.Tech()
	}
	a.History = append(a.History, h)
	a = generateElapsed(a)
	return a
}

func generateElapsed(a Asset) Asset {
	l := len(a.History) - 1
	a.History[l].Elapsed = time.Since(a.History[l-1].Start)
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
			r := 0
			if l > 1 {
				r = rand.Intn(l)
			}
			a := Assets[r]
			if a.Status == "completed" {
				removeAsset(r)
			} else {
				switch a.Status {
				case "pending":
					a.Status = "processing"
					a = generateHistory("processing", a)
				case "error":
					q := rand.Intn(10)
					if q < 2 {
						a.Status = "processing"
						a = generateHistory("processing", a)
					} else {
						a = generateHistory("error", a)
					}
				case "processing":
					q := rand.Intn(10)
					if q < 2 {
						a.Status = "error"
						a = generateHistory("error", a)
					} else {
						a.Status = "completed"
						a = generateHistory("completed", a)
					}
				default:
					// do nothing
				}
				Assets[r] = a
				a.Update = true
				l := len(a.History) - 1
				a.Elapsed = a.History[l].Start.Sub(a.History[0].Start)
				if a.ID == monitorID {
					a.Monitor = a.ID
				}
				updateCh <- a
			}
		}
	}

	ticker.Stop()

	close(updateCh)

	// color.Println(color.Yellow("generateUpdates: Completed"))
}
