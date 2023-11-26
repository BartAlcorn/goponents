package todos

import (
	"fmt"
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

	web.Respond(w, r, i, "pkg/todos/tmpls/*.gohtml", "grid")
}
