package todos

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Delete remove a ToDo by ID
func Delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	item, err := DeleteById(idParam)
	if err != nil {
		fmt.Println("error GetById", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("DELETIN", idParam, item.DeletedCount)
	w.WriteHeader(204)

}
