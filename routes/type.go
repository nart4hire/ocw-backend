package routes

import "github.com/go-chi/chi/v5"

type RouteCollection interface {
	Register() []RouteGroup
}

type RouteGroup interface {
	Register(chi.Router)
}
