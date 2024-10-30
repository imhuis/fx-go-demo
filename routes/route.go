package routes

import (
	"go.uber.org/fx"
	"net/http"
)

type Route interface {
	http.Handler

	Pattern() string
}

// AsRoute annotates the given constructor to state that
// it provides a route to the "routes" group.
func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
