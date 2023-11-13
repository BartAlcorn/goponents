package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/bartalcorn/goponents/web/todos"
)

// Update calls the UpdateById func
func Update(w http.ResponseWriter, r *http.Request) {
	var body *todos.Todo

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Println("error decoding body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idParam := chi.URLParam(r, "id")

	_, err := todos.UpdateById(idParam, body)
	if err != nil {
		fmt.Println("error UpdateById", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	item, err := todos.GetById(idParam)
	if err != nil {
		fmt.Println("error GetById", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t, err := template.ParseGlob("web/todos/tmpls/*.gohtml")
	if err != nil {
		fmt.Println("error parsing gohtml", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = t.Execute(w, item)
	if err != nil {
		fmt.Println("error executing", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

}
