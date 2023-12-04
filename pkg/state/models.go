package state

type state struct {
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

var State state

func init() {
	State = state{
		Module: "home",
		Modules: []WebModule{
			{Name: "home", Icon: "home", Path: "home"},
			{Name: "todos", Icon: "todos", Path: "todos"},
			{Name: "sse", Icon: "sse1", Path: "sse"},
			{Name: "min", Icon: "sse2", Path: "min"},
		},
		Theme:  "discolight",
		Themes: []string{"light", "dark", "cupcake", "garden", "retro", "business", "aqua", "nord", "discolight"},
	}
	// auto-parse and include tailwind themes from config
	// State.Themes = GetThemes()
}
