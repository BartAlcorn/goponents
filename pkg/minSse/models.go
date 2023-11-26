package minsse

import (
	"time"
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

type History struct {
	Process     string        `json:"process"`
	Status      string        `json:"status"`
	Description string        `json:"description"`
	Start       time.Time     `json:"start"`
	End         time.Time     `json:"end,omitempty"`
	Elapsed     time.Duration `json:"duration"`
}
