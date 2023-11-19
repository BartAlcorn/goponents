package ssesimulator

import (
	"fmt"
	"html/template"
)

var t *template.Template
var Assets []Asset

func init() {
	// Parse the template and store it
	tmpl, err := template.ParseFiles("pkg/ssesimulator/tmpls/simulatorLine.gohtml", "pkg/ssesimulator/tmpls/statusicons.gohtml")
	if err != nil {
		fmt.Println("error parsing gohtml", err)
	}
	t = tmpl

	Assets = make([]Asset, 0)
}
