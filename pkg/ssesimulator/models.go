package ssesimulator

import "time"

type Asset struct {
	ID      string        `json:"id"`
	Title   string        `json:"title"`
	Status  string        `json:"Status"`
	History []History     `json:"history"`
	Elapsed time.Duration `json:"elapsed"`
	Update  bool          `json:"-"`
	Monitor string        `json:"-"`
}

type History struct {
	Process     string        `json:"process"`
	Status      string        `json:"status"`
	Description string        `json:"description"`
	Start       time.Time     `json:"start"`
	End         time.Time     `json:"end,omitempty"`
	Elapsed     time.Duration `json:"duration"`
}
