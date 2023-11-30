package minsse

import (
	"fmt"
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
	Counts  StatBlock     `json:"-"` // transport for HeroCard block, do NOT store in DB
	Metrics StatBlock     `json:"-"` // transport for HeroCard block, do NOT store in DB
}

type History struct {
	Process     string        `json:"process"`
	Status      string        `json:"status"`
	Description string        `json:"description"`
	Start       time.Time     `json:"start"`
	End         time.Time     `json:"end,omitempty"`
	Elapsed     time.Duration `json:"duration"`
}

type StatBlock struct {
	Name        string `json:"name"`        // name of this block
	Total       int    `json:"total"`       // total numbers of events processed
	Errors      int    `json:"errors"`      // number of in error status
	ErrorP      string `json:"errorp"`      // percentage of errors
	Failures    int    `json:"failures"`    // number of in failure status
	FailureP    string `json:"failurep"`    // percentage of failures
	Processings int    `json:"processings"` // number of in processing status
	ProcessingP string `json:"processingp"` // Ppercentage of processings
	Pendings    int    `json:"pendings"`    // number of in pending status
	PendingP    string `json:"pendingp"`    // percentage of pendings
	Completeds  int    `json:"completeds"`  // number of in completed status
	CompletedP  string `json:"completedp"`  // percentage of completeds
	Others      int    `json:"others"`      // number of in other status
	OtherP      string `json:"othersp"`     // percentage of others
}

func (t *StatBlock) Compute() {
	var vt float64 = float64(t.Total)
	var ve float64 = float64(t.Errors)
	var vf float64 = float64(t.Failures)
	var vs float64 = float64(t.Processings)
	var vp float64 = float64(t.Pendings)
	var vc float64 = float64(t.Completeds)
	var vo float64 = float64(t.Others)
	t.ErrorP = fmt.Sprintf("%.1f", (ve/vt)*100)
	t.FailureP = fmt.Sprintf("%.1f", (vf/vt)*100)
	t.ProcessingP = fmt.Sprintf("%.1f", (vs/vt)*100)
	t.PendingP = fmt.Sprintf("%.1f", (vp/vt)*100)
	t.CompletedP = fmt.Sprintf("%.1f", (vc/vt)*100)
	t.OtherP = fmt.Sprintf("%.1f", (vo/vt)*100)
}
