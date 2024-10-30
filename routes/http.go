package routes

import (
	"context"
	"go.uber.org/fx"
	"net/http"
)

func NewHttpServer(lc fx.Lifecycle, mux *http.ServeMux) *http.Server {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go srv.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

// ==================================================

// HttpServerParams 1. define a struct with fx.In tag
type HttpServerParams struct {
	fx.In

	HttpRoutes []Route `group:"routes"`
}

func NewServeMux1(params HttpServerParams) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range params.HttpRoutes {
		mux.Handle(route.Pattern(), route)
	}
	return mux
}

// ==================================================

// NewServeMux 2. directly put Route type in the function signature
func NewServeMux(routes []Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.Handle(route.Pattern(), route)
	}
	return mux
}
