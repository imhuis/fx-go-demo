package main

import (
	"github.azc.ext.hp.com/jia-hui-ye/fx-demo/config"
	"github.azc.ext.hp.com/jia-hui-ye/fx-demo/routes"
	"go.uber.org/fx"
	"net/http"
)

func main() {
	var _ = fx.Module("server", fx.Supply(&config.AppConfig{}))

	fx.New(
		fx.Provide(
			routes.NewHttpServer,
			fx.Annotate(
				routes.NewServeMux,
				fx.ParamTags(`group:"routes"`),
			),
			// Fx does not allow two instances of the same type to be present in the container without annotating them.
			fx.Annotate(routes.NewEchoHandler, fx.As(new(routes.Route)), fx.ResultTags(`group:"routes"`)),
			routes.AsRoute(routes.NewEchoTextHandler),
		),
		fx.Invoke(
			// Invoke case:
			// 1. Starting a background worker
			// 2. Configuring a global logger

			func(s *http.Server) {},
		),
	).Run()
}
