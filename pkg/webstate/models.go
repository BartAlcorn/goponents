package webstate

type WebState struct {
	Module  string
	Modules []WebModule
	Theme   string
	Themes  []string
}

type WebModule struct {
	Name string
	Icon string
}

var State WebState

func init() {
	State = WebState{
		Module: "home",
		Modules: []WebModule{
			{Name: "home", Icon: "icons-home"},
			{Name: "todos", Icon: ""},
			{Name: "sse", Icon: ""},
			{Name: "min", Icon: ""},
			{Name: "acquire", Icon: ""},
		},
		Theme:  "cupcake",
		Themes: []string{"light", "dark", "cupcake", "garden", "retro", "business", "aqua"},
	}
}
