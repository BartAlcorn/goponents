package minsse

import (
	"time"

	"github.com/google/uuid"
)

func MakeAsset() Asset {
	var h History
	h.Process = "Receiving"
	h.Status = "pending"
	h.Start = time.Now()
	h.Elapsed = (0 * time.Millisecond)

	var a Asset
	a.ID = uuid.NewString()
	a.Title = FakeMovie()
	a.Status = "pending"
	a.History = append(a.History, h)
	return a
}
