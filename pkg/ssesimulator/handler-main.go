package ssesimulator

import (
	"fmt"
	"html/template"
	"net/http"
)

// Create calls the Insert func
func ShowMain(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseGlob("pkg/ssesimulator/tmpls/*.gohtml")
	if err != nil {
		fmt.Println("error template parsing handle_main", err)
	}

	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println("error executing", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
