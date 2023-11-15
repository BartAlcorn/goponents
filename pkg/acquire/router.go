package acquire

import "github.com/go-chi/chi/v5"

func Routes(r chi.Router) {
	r.Get("/", Index)
	r.Get("/{id}", Read)
	r.Get("/brief/{id}", BriefDetails)
	r.Get("/details/{id}", StatusDetails)
	r.Get("/modalbtn", ModalBtn)
}
