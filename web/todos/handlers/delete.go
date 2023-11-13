package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/bartalcorn/goponents/web/todos"
)

// Delete remove a ToDo by ID
func Delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	item, err := todos.DeleteById(idParam)
	if err != nil {
		fmt.Println("error GetById", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("DELETE", idParam, item.DeletedCount)
	w.WriteHeader(204)

}
