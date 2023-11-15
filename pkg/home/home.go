package web

import (
	"fmt"
	"html/template"
	"net/http"

	web "github.com/bartalcorn/goponents/web/state"
)

func Home(w http.ResponseWriter, r *http.Request) {

	web.State.Module = "home"

	t, err := template.ParseFiles("pkg/home/home.gohtml", "web/index/tech.gohtml", "web/index/why.gohtml", "web/index/stack.gohtml")
	if err != nil {
		fmt.Println("error parsing gohtml", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = t.Execute(w, &web.State)
	if err != nil {
		fmt.Println("error executing", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
