package ssesimulator

import (
	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {
	r.Get("/", ShowMain)
	r.Get("/events", Simulate)
	r.Get("/start", HandleStart)
	r.Get("/stop", HandleStop)
	r.Get("/details/{id}", HandleDetails)
}
