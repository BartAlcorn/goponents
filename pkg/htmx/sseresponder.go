package htmx

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/Masterminds/sprig/v3"
)

// SSERespond parses the template(s) and return an SSE formatted data block
func SSERespond(item any, t string, block string, event string) (string, error) {
	var tpl bytes.Buffer
	tmpl, err := template.New("main.gohtml").Funcs(sprig.FuncMap()).ParseGlob(t)
	if err != nil {
		fmt.Println("error parsing gohtml", err)
	}
	if block == "" {
		err = tmpl.Execute(&tpl, item)
	} else {
		err = tmpl.ExecuteTemplate(&tpl, block, item)
	}
	if err != nil {
		fmt.Println("error executing", err)
	}

	sb := strings.Builder{}
	sanitized := sanitize(tpl.String())
	sb.WriteString(fmt.Sprintf("event: %s\n", event))
	sb.WriteString(fmt.Sprintf("data: %v\n\n", sanitized))

	fmt.Println(sb.String())

	return sb.String(), nil
}
