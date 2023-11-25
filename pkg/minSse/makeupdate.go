package minsse

import (
	"context"
	"math/rand"
)

func MakeUpdate(ctx context.Context, updateCh chan<- Asset) {
	r := 0
	l := len(Assets)
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
			a.generateHistory("processing")
		case "error":
			q := rand.Intn(100)
			if q < 34 {
				a.Status = "processing"
				a.generateHistory("processing")
			} else {
				a.generateHistory("error")
			}
		case "processing":
			q := rand.Intn(100)
			if q < 34 {
				a.Status = "error"
				a.generateHistory("error")
			} else {
				a.Status = "completed"
				a.generateHistory("completed")
				a.complete()
			}
		default:
			// do nothing
		}
		Assets[r] = a
		a.calcElapsed()
		if a.ID == monitorID {
			a.Monitor = a.ID
		}
		a.Update = true
		updateCh <- a
	}
}
