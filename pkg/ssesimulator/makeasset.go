package ssesimulator

import (
	"time"

	"github.com/google/uuid"
)

func MakeAsset() Asset {
	var h History
	h.Process = "Receiving"
	h.Status = "pending"
	h.Start = time.Now()

	var a Asset
	a.ID = uuid.NewString()
	a.Title = FakeMovie()
	a.Status = "completed"
	a.History = append(a.History, h)
	return a
}
