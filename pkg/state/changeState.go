package state

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ChangeState(w http.ResponseWriter, r *http.Request) {
	s := chi.URLParam(r, "theme")
	State.Theme = s
	_, err := w.Write([]byte(s))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}
