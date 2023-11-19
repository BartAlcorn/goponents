package ssesimulator

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/labstack/gommon/color"
)

func generateHistory(status string, a Asset) Asset {
	var h History
	h.Status = status
	h.Start = time.Now()
	a.History = append(a.History, h)
	return a
}

func removeAsset(i int) {
	Assets = append(Assets[:i], Assets[i+1:]...)
	fmt.Println("new len of Assets", len(Assets))
}

func generateUpdates(ctx context.Context, updateCh chan<- Asset) {
	ticker := time.NewTicker(500 * time.Millisecond)
	color.Println(color.Green("generateUpdates: started"))

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
			fmt.Printf("ASSETS: %v of %v\n", r, len(Assets))
			a := Assets[r]
			if a.Status == "completed" {
				fmt.Println("Already completed", a.ID)
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
				fmt.Println("Updating ", a.ID, a.Status)
				updateCh <- a
			}
		}
	}

	ticker.Stop()

	close(updateCh)

	color.Println(color.Yellow("generateUpdates: Completed"))

}
