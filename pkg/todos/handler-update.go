package todos

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/bartalcorn/goponents/pkg/web"
	"github.com/go-chi/chi/v5"
)

// Update calls the UpdateById func
func Update(w http.ResponseWriter, r *http.Request) {
	var body *Todo

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Println("error decoding body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idParam := chi.URLParam(r, "id")

	_, err := UpdateById(idParam, body)
	if err != nil {
		fmt.Println("error UpdateById", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	item, err := GetById(idParam)
	if err != nil {
		fmt.Println("error GetById", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t, err := template.ParseGlob("pkg/todos/tmpls/*.gohtml")
	if err != nil {
		fmt.Println("error parsing gohtml", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	web.Respond(w, r, item, t)
}
