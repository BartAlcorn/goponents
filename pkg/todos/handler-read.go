package todos

import (
	"fmt"
	"html/template"
	"net/http"

	state "github.com/bartalcorn/goponents/pkg/state"
	"github.com/bartalcorn/goponents/pkg/web"
)

// Read calls the GetAll func
func Read(w http.ResponseWriter, r *http.Request) {
	state.State.Module = "todos"

	i, err := GetAll()
	if err != nil {
		fmt.Println("error getting all", err)
	}

	t, err := template.ParseFiles("pkg/todos/tmpls/todos-grid.gohtml", "pkg/todos/tmpls/todos-form.gohtml")
	web.Respond(w, r, i, t)

}
