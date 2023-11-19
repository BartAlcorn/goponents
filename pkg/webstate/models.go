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
	Path string
}

var State WebState

func init() {
	State = WebState{
		Module: "home",
		Modules: []WebModule{
			{Name: "home", Icon: "home", Path: "home"},
			{Name: "todos", Icon: "todos", Path: "todos"},
			{Name: "sse", Icon: "sse1", Path: "sse"},
			{Name: "min", Icon: "sse2", Path: "min"},
		},
		Theme:  "cupcake",
		Themes: []string{"light", "dark", "cupcake", "garden", "retro", "business", "aqua"},
	}
}
