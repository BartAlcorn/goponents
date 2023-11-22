package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	home "github.com/bartalcorn/goponents/pkg/home"
	idx "github.com/bartalcorn/goponents/pkg/index"
	sse "github.com/bartalcorn/goponents/pkg/minSse"
	todos "github.com/bartalcorn/goponents/pkg/todos"
	webstate "github.com/bartalcorn/goponents/pkg/webstate"
)

func (a *AppConfig) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "HX-Request"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Get("/", idx.Index)

	//handle static files like CSS and the few, but necessary JS files
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// grouped routes
	router.Route("/home/", home.Routes)
	router.Route("/events", a.loadSSERoutes)
	router.Route("/min", sse.Routes)
	router.Route("/sse", a.loadSSERoutes)
	router.Route("/state", webstate.Routes)
	router.Route("/todos", todos.Routes)

	a.router = router
}

// Simple Server Sent Events demonstrator
func (a *AppConfig) loadSSERoutes(router chi.Router) {
	router.Get("/", idx.SimpleSseExample)
	router.Get("/events", minSse)
}
