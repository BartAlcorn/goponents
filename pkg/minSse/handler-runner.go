package minSse

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Masterminds/sprig/v3"
	"github.com/go-chi/chi/v5"
)

func HandleStart(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("details.gohtml").Funcs(sprig.FuncMap()).ParseGlob("pkg/minSse/tmpls/*.gohtml")
	if err != nil {
		fmt.Println("error parsin template formatAsset", err)
	}

	err = t.ExecuteTemplate(w, "sseStart", nil)
	if err != nil {
		fmt.Println("error executing", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func HandleStop(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("details.gohtml").Funcs(sprig.FuncMap()).ParseGlob("pkg/minSse/tmpls/*.gohtml")
	if err != nil {
		fmt.Println("error parsin template formatAsset", err)
	}

	err = t.ExecuteTemplate(w, "sseStop", nil)
	if err != nil {
		fmt.Println("error executing", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func HandleDetails(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("details.gohtml").Funcs(sprig.FuncMap()).ParseGlob("pkg/minSse/tmpls/*.gohtml")
	if err != nil {
		fmt.Println("error parsin template formatAsset", err)
	}

	idParam := chi.URLParam(r, "id")
	monitorID = idParam

	var a Asset
	for _, v := range Assets {
		if v.ID == idParam {
			a = v
		}
	}

	err = t.ExecuteTemplate(w, "details", a)
	if err != nil {
		fmt.Println("error executing", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
