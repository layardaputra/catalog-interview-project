package app

import (
	"net/http"

	"github.com/go-chi/chi"
)

// RegisterProductRoutes registers API routes.
func RegisterRoutes(app *App, router *chi.Mux) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the API"))
	})

	router.Group(func(r chi.Router) {
		r.Route("/api/v1", func(r chi.Router) {
			r.Get("/products/{productID}", app.ProductHandler.GetProduct)
			r.Get("/products", app.ProductHandler.ListProduct)
			r.Post("/products", app.ProductHandler.CreateProduct)
			r.Put("/products/{productID}", app.ProductHandler.UpdateProduct)

			// review
			r.Post("/products/{productID}/reviews", app.ProductHandler.AddProductReview)
		})
	})
}
