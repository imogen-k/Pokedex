package handlers

import (
	"net/http"
)

type GetPokemonInformationHandler struct {
}

func NewGetPokemonInformationHandler() *GetPokemonInformationHandler {
	return &GetPokemonInformationHandler{}
}

// Handle GET /pokemon/{name}
func (p *GetPokemonInformationHandler) Handle(w http.ResponseWriter, r *http.Request) {
}
