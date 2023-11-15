package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/bartalcorn/goponents/pkg/todos"
	web "github.com/bartalcorn/goponents/pkg/webstate"
)

// Read calls the GetAll func
func Read(w http.ResponseWriter, r *http.Request) {
	web.State.Module = "todos"
	items, err := todos.GetAll()
	if err != nil {
		fmt.Println("error getting all", err)
	}

	// if not HTMX, return JSON
	isHTMX := r.Header.Get("HX-Request")
	if isHTMX != "true" {
		res, err := json.Marshal(items)
		if err != nil {
			fmt.Println("failed to marshal:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
		_, err = w.Write(res)
		if err != nil {
			fmt.Println("failed to write", err)
		}
		return
	}

	// else return HTMX
	t, err := template.ParseFiles("pkg/todos/tmpls/todos-grid.gohtml", "pkg/todos/tmpls/todos-form.gohtml")
	if err != nil {
		fmt.Println("error parsing gohtml", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = t.Execute(w, items)
	if err != nil {
		fmt.Println("error executing", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
