package server

import (
	"Pokedex/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
)

func NewChiRouter() *chi.Mux {
	port := "8080"

	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	log.Printf("Starting up on http://localhost:%s", port)
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	getPokemonInformationHandler := handlers.NewGetPokemonInformationHandler()
	getTranslatedPokemonDescription := handlers.NewGetTranslatedPokemonDescriptionHandler()

	r.Route("/pokemon", func(r chi.Router) {
		r.Get("/{name}", getPokemonInformationHandler.Handle)
		r.Get("/{name}/", getPokemonInformationHandler.Handle)

		r.Get("/translated/{name}", getTranslatedPokemonDescription.Handle)
		r.Get("/translated/{name}/", getTranslatedPokemonDescription.Handle)

	})

	log.Fatal(http.ListenAndServe(":"+port, r))

	return r
}
