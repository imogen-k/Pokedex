package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io/ioutil"
	"net/http"
	"strings"
)

type GetPokemonInformationHandler struct {
}

func NewGetPokemonInformationHandler() *GetPokemonInformationHandler {
	return &GetPokemonInformationHandler{}
}

type PokemonSpecies struct {
	Name    string `json:"name"`
	Species struct {
		URL string `json:"url"`
	}
}

type PokemonInformation struct {
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Name string `json:"name"`
		} `json:"language"`
	} `json:"flavor_text_entries"`
	Habitat struct {
		Name string `json:"name"`
	} `json:"habitat"`
	IsLegendary bool   `json:"is_legendary"`
	Name        string `json:"name"`
}

type GetPokemonInformationResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Habitat     string `json:"habitat"`
	IsLegendary bool   `json:"is_legendary"`
}

// Handle GET /pokemon/{name}
func (p *GetPokemonInformationHandler) Handle(w http.ResponseWriter, r *http.Request) {
	name := chi.RouteContext(r.Context()).URLParam("name")

	// find pokemon species URL with pokemon name
	url, err := GetSpeciesUrlFromName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// use url to call API to find habitat, description and legendary status
	pokemonInformation, err := GetPokemonInformationFromUrl(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get english description from flavor text entries
	englishDescription := GetEnglishDescription(pokemonInformation)

	getPokemonInformationResponse := &GetPokemonInformationResponse{
		Name:        name,
		Description: englishDescription,
		Habitat:     pokemonInformation.Habitat.Name,
		IsLegendary: pokemonInformation.IsLegendary,
	}

	jsonResponse, err := json.Marshal(getPokemonInformationResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func GetSpeciesUrlFromName(name string) (string, error) {
	response, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name))
	if err != nil {
		return "", err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	var urlResponseObject PokemonSpecies
	err = json.Unmarshal(responseData, &urlResponseObject)
	if err != nil {
		return "", err
	}

	return urlResponseObject.Species.URL, nil
}

func GetPokemonInformationFromUrl(url string) (PokemonInformation, error) {
	response, err := http.Get(url)
	if err != nil {
		return PokemonInformation{}, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return PokemonInformation{}, err
	}

	defer response.Body.Close()

	var informationResponseObject PokemonInformation
	err = json.Unmarshal(responseData, &informationResponseObject)
	if err != nil {
		return PokemonInformation{}, err
	}

	return informationResponseObject, nil
}

func GetEnglishDescription(information PokemonInformation) string {
	var englishDescription string
	for _, description := range information.FlavorTextEntries {
		if description.Language.Name == "en" {
			englishDescription = description.FlavorText
			englishDescription = strings.Replace(englishDescription, "\n", " ", -1)

		}
	}
	return englishDescription
}
