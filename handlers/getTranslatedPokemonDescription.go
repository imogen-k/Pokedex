package handlers

import (
	"net/http"
)

type GetTranslatedPokemonDescriptionHandler struct {
}

func NewGetTranslatedPokemonDescriptionHandler() *GetTranslatedPokemonDescriptionHandler {
	return &GetTranslatedPokemonDescriptionHandler{}
}

// Handle GET /pokemon/translated/{name}
func (p *GetTranslatedPokemonDescriptionHandler) Handle(w http.ResponseWriter, r *http.Request) {
}
