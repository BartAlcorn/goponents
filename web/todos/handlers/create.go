package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/google/uuid"

	"github.com/bartalcorn/goponents/web/todos"
)

// Create calls the Insert func
func Create(w http.ResponseWriter, r *http.Request) {
	item := todos.Todo{
		ID:   uuid.NewString(),
		Task: r.PostFormValue("task"),
		// Done: body.Done,
	}

	// Call repo
	err := todos.Insert(item)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t, err := template.ParseGlob("web/todos/tmpls/*.gohtml")
	if err != nil {
		fmt.Println("error parsing gohtml", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = t.ExecuteTemplate(w, "todo-list-element", item)
	if err != nil {
		fmt.Println("error executing", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusAccepted)
}
