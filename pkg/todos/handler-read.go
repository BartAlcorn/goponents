package todos

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/bartalcorn/goponents/pkg/web"
	state "github.com/bartalcorn/goponents/pkg/webstate"
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
