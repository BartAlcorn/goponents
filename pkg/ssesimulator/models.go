package ssesimulator

import "time"

type Asset struct {
	ID      string    `json:"id"`
	Title   string    `json:"title"`
	Status  string    `json:"Status"`
	History []History `json:"history"`
	Update  bool      `json:"update"`
}

type History struct {
	Process string        `json:"process"`
	Status  string        `json:"status"`
	Start   time.Time     `json:"start"`
	End     time.Time     `json:"end,omitempty"`
	Elasped time.Duration `json:"duration,omitempty"`
}
