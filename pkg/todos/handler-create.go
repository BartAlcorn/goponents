package todos

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/bartalcorn/goponents/pkg/web"
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

	t, err := template.ParseGlob("pkg/todos/tmpls/*.gohtml")
	if err != nil {
		fmt.Println("error parsing gohtml", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	web.RespondWithTemplate(w, r, item, t, "todo-list-element")
}
