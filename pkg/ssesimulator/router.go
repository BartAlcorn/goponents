package ssesimulator

import (
	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {
	r.Get("/", ShowMain)
	r.Get("/events", Simulate)
}