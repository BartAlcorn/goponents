package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/bartalcorn/goponents/web/todos"
)

// Update calls the UpdateById func
func ToggleDone(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	var item *todos.Todo

	item, err := todos.GetById(idParam)
	if err != nil {
		fmt.Println("error GetById", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	item.Done = !item.Done

	_, err = todos.UpdateById(idParam, item)
	if err != nil {
		fmt.Println("error UpdateById", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)

}
