package application

import (
	"net/http"

	// sse "github.com/alexandrevicenzi/go-sse"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	home "github.com/bartalcorn/goponents/pkg/home"
	idx "github.com/bartalcorn/goponents/pkg/index"
	orders "github.com/bartalcorn/goponents/pkg/orders"
	"github.com/bartalcorn/goponents/pkg/ssesimulator"
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

	// SSE Server
	// s := sse.NewServer(nil)
	// defer s.Shutdown()

	router.Get("/", idx.Index)

	//handle static files like CSS and the few, but necessary JS files
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// grouped routes
	router.Route("/home/", home.Routes)
	router.Route("/orders", a.loadOrderRoutes)
	router.Route("/events", a.loadSSERoutes)
	router.Route("/state", webstate.Routes)
	router.Route("/todos", todos.Routes)
	router.Route("/sse", a.loadSSERoutes)
	router.Route("/min", ssesimulator.Routes)

	a.router = router
}

// Server Sent Events demonstrator
func (a *AppConfig) loadSSERoutes(router chi.Router) {
	router.Get("/", idx.ServerSentEvents)
	router.Get("/events", SSESimulator)
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
