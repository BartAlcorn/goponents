package minsse

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/Masterminds/sprig/v3"
)

func formatReturn(event string, data any, tmpl string) (string, error) {
	var tpl bytes.Buffer
	t, err := template.New("simulator.gohtml").Funcs(sprig.FuncMap()).ParseGlob("pkg/minSse/tmpls/*.gohtml")
	if err != nil {
		fmt.Println("error parsin template formatReturn", err)
	}
	err = t.ExecuteTemplate(&tpl, tmpl, data)
	if err != nil {
		fmt.Println("error executing", err)
		// w.WriteHeader(http.StatusInternalServerError)
	}

	sb := strings.Builder{}
	sanitized := strings.ReplaceAll(tpl.String(), "\n  ", "")
	sanitized = strings.ReplaceAll(sanitized, "\n", "")
	sb.WriteString(fmt.Sprintf("event: %s\n", event))
	sb.WriteString(fmt.Sprintf("data: %v\n\n", sanitized))

	return sb.String(), nil
}
