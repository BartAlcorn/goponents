package todos

import (
	"net/http"

	"github.com/bartalcorn/goponents/pkg/htmx"
	"github.com/google/uuid"
)

// Create calls the Insert func
func Create(w http.ResponseWriter, r *http.Request) {
	item := Todo{
		ID:   uuid.NewString(),
		Task: r.PostFormValue("task"),
	}

	err := Insert(item)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	htmx.Respond(w, r, item, "pkg/todos/tmpls/*.gohtml", "todo-list-element")
}
