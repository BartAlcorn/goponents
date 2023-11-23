package simplesse

import (
	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {
	r.Get("/", SimpleSseExample)
	r.Get("/events", SimpleSse)
}
