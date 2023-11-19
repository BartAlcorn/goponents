package ssesimulator

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Masterminds/sprig/v3"
	"github.com/go-chi/chi/v5"
)

func HandleStart(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("pkg/ssesimulator/tmpls-partials/startsse.gohtml")
	if err != nil {
		fmt.Println("error template parsing HandleStart", err)
	}

	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println("error executing", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func HandleStop(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("pkg/ssesimulator/tmpls-partials/stopsse.gohtml")
	if err != nil {
		fmt.Println("error template parsing HandleStart", err)
	}

	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println("error executing", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func HandleDetails(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("details.gohtml").Funcs(sprig.FuncMap()).ParseGlob("pkg/ssesimulator/tmpls-partials/details.gohtml")
	if err != nil {
		fmt.Println("error parsin template formatAsset", err)
	}

	idParam := chi.URLParam(r, "id")

	var a Asset
	for _, v := range Assets {
		if v.ID == idParam {
			a = v
		}
	}

	err = t.Execute(w, a)
	if err != nil {
		fmt.Println("error executing", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
