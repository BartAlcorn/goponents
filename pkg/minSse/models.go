package minsse

import (
	"math/rand"
	"time"

	"github.com/bartalcorn/excuses"
)

type Asset struct {
	ID      string        `json:"id"`
	Title   string        `json:"title"`
	Status  string        `json:"Status"`
	History []History     `json:"history"`
	Elapsed time.Duration `json:"elapsed"`
	Update  bool          `json:"-"` // flags an asset to be updated, do NOT store in DB
	Monitor string        `json:"-"` // flags an asset to be broadcast on Update, do NOT store in DB
}

func (a *Asset) generateHistory(status string) {
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
	a.generateElapsed()
}

func (a *Asset) generateElapsed() {
	l := len(a.History)
	if l == 0 {
		return
	}
	a.History[1].End = time.Now()
	a.History[1].Elapsed = a.History[1].End.Sub(a.History[1].Start).Round(3 * time.Millisecond)
}

func (a *Asset) calcElapsed() {
	a.Elapsed = 0 * time.Millisecond
	for _, v := range a.History {
		a.Elapsed += v.Elapsed
	}
}

func (a *Asset) complete() {
	a.History[0].End = time.Now()
	a.generateElapsed()
	a.Elapsed = 0 * time.Millisecond
	for _, v := range a.History {
		a.Elapsed += v.Elapsed
	}
}

type History struct {
	Process     string        `json:"process"`
	Status      string        `json:"status"`
	Description string        `json:"description"`
	Start       time.Time     `json:"start"`
	End         time.Time     `json:"end,omitempty"`
	Elapsed     time.Duration `json:"duration"`
}
