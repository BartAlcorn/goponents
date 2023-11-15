package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	orders "github.com/bartalcorn/goponents/pkg/orders"
	acquire "github.com/bartalcorn/goponents/web/acquire"
	appconfig "github.com/bartalcorn/goponents/web/appconfig"
	home "github.com/bartalcorn/goponents/web/home"
	idx "github.com/bartalcorn/goponents/web/index"
	state "github.com/bartalcorn/goponents/web/state"
	todos "github.com/bartalcorn/goponents/web/todos/handlers"
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
	router.Get("/sse/", idx.ServerSentEvents)

	//handle static files like CSS and the few, but necessary JS files
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// grouped routes
	router.Route("/home/", a.loadHomeRoutes)
	router.Route("/acquire", a.loadAcquireRoutes)
	router.Route("/appconfigs", a.loadAppConfigRoutes)
	router.Route("/orders", a.loadOrderRoutes)
	router.Route("/events", a.loadSSERoutes)
	router.Route("/state", a.loadStateRoutes)
	router.Route("/todos", a.loadToDosRoutes)

	a.router = router
}

// Home
func (a *AppConfig) loadHomeRoutes(router chi.Router) {
	router.Get("/", home.Home)
}

// Acquire
func (a *AppConfig) loadAcquireRoutes(router chi.Router) {
	router.Get("/", acquire.Index)
	router.Get("/{id}", acquire.Read)
	router.Get("/brief/{id}", acquire.BriefDetails)
	router.Get("/details/{id}", acquire.StatusDetails)
	router.Get("/modalbtn", acquire.ModalBtn)
}

// AppConfigs
func (a *AppConfig) loadAppConfigRoutes(router chi.Router) {
	router.Get("/", appconfig.Read)
}

// ToDos
func (a *AppConfig) loadToDosRoutes(router chi.Router) {
	router.Post("/", todos.Create)
	router.Get("/", todos.Read)
	router.Put("/{id}", todos.Update)
	router.Put("/{id}/done", todos.ToggleDone)
	router.Delete("/{id}", todos.Delete)
}

// State
func (a *AppConfig) loadStateRoutes(router chi.Router) {
	router.Get("/theme/{theme}", state.ChangeState)
}

// Server Sent Events demonstrator
func (a *AppConfig) loadSSERoutes(router chi.Router) {
	router.Get("/", SSESimulator)
}

// Orders (JSON return only! NOT HTMX enabled)
func (a *AppConfig) loadOrderRoutes(router chi.Router) {
	orderHandler := &orders.OrderHandler{
		Repo: &orders.RedisRepo{
			Client: a.RedisDB,
		},
	}

	router.Post("/", orderHandler.Create)
	router.Get("/", orderHandler.List)
	router.Get("/{id}", orderHandler.GetByID)
	router.Put("/{id}", orderHandler.UpdateByID)
	router.Delete("/{id}", orderHandler.DeleteByID)
}
