package todos

import (
	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {
	r.Post("/", Create)
	r.Get("/", Read)
	r.Put("/{id}", Update)
	r.Put("/{id}/done", ToggleDone)
	r.Delete("/{id}", Delete)
}
