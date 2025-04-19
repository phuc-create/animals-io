package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/phuc-create/animals-io/internal/api/v1/handlers"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/users", handlers.UserRoutes())
		r.Mount("/auth", handlers.AuthenticationHandlers())
	})
	return r
}
