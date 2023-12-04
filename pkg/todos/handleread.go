package todos

import (
	"fmt"
	"net/http"

	"github.com/bartalcorn/goponents/pkg/htmx"
	state "github.com/bartalcorn/goponents/pkg/state"
)

// Read calls the GetAll func
func Read(w http.ResponseWriter, r *http.Request) {
	state.State.Module = "todos"

	i, err := GetAll()
	if err != nil {
		fmt.Println("error getting all", err)
	}

	htmx.Respond(w, r, i, "pkg/todos/tmpls/*.gohtml", "main")
}
