package router

import (
	"app/config"
	"app/controller"
	"app/middlewares"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
)

func Router() http.Handler {
	app := chi.NewRouter()

	app.Use(middleware.RequestID)
	app.Use(middleware.RealIP)
	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	app.Use(cors.Handler)

	middlewares := middlewares.NewMiddlewares()
	shopController := controller.NewShopController()

	app.Route("/shop/api/v1", func(r chi.Router) {
		r.Route("/public", func(public chi.Router) {
		})

		r.Route("/protected", func(protected chi.Router) {
			protected.Use(jwtauth.Verifier(config.GetJWT()))
			protected.Use(jwtauth.Authenticator(config.GetJWT()))
			protected.Use(middlewares.ValidateExpAccessToken())

			protected.Route("/shop", func(shop chi.Router) {
				shop.Get("/check-duplicate", shopController.CheckDuplicateShop)
				shop.Get("/", shopController.GetShop)
				shop.Get("/type-product", shopController.GetTypeProduct)

				shop.Post("/", shopController.CreateShop)
				shop.Post("/type-product", shopController.CreateTypeProduct)
			})
		})
	})

	return app
}
