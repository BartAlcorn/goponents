package acquire

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
)

// Read calls the GetAll func
func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseGlob("web/acquire/tmpls/acquireIndex.gohtml")
	if err != nil {
		fmt.Println("error parsing gohtml", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println("error executing", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ModalBtn(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseGlob("web/acquire/tmpls/acquireModalDetails.gohtml")
	if err != nil {
		fmt.Println("error parsing gohtml", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println("error executing", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func StatusDetails(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/acquire/tmpls/acquireStatusIcon.gohtml", "web/acquire/tmpls/statusicons.gohtml")
	if err != nil {
		fmt.Println("error parsing gohtml", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println("error executing", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func BriefDetails(w http.ResponseWriter, r *http.Request) {
	item, err := Find(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("error getting all", err)
	}

	t, err := template.ParseFiles("web/acquire/tmpls/acquireBriefDetails.gohtml", "web/acquire/tmpls/statusicons.gohtml")
	if err != nil {
		fmt.Println("error parsing gohtml", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = t.Execute(w, item)
	if err != nil {
		fmt.Println("error executing", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Read calls the GetAll func
func Read(w http.ResponseWriter, r *http.Request) {
	item, err := Find(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("error getting all", err)
	}

	// if not HTMX, return JSON
	isHTMX := r.Header.Get("HX-Request")
	if isHTMX != "true" {
		res, err := json.Marshal(item)
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
	t, err := template.ParseFiles("web/acquire/tmpls/acquireDetails.gohtml", "web/acquire/tmpls/statusicons.gohtml")
	if err != nil {
		fmt.Println("error parsing gohtml", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = t.Execute(w, item)
	if err != nil {
		fmt.Println("error executing", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
