package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type GetTranslatedPokemonDescriptionHandler struct {
}

func NewGetTranslatedPokemonDescriptionHandler() *GetTranslatedPokemonDescriptionHandler {
	return &GetTranslatedPokemonDescriptionHandler{}
}

type TranslatedDescriptionResponse struct {
	Success struct {
		Total int `json:"total"`
	} `json:"success"`
	Contents struct {
		Translated string `json:"translated"`
	} `json:"contents"`
}

// Handle GET /pokemon/translated/{name}
func (p *GetTranslatedPokemonDescriptionHandler) Handle(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/pokemon/translated/")

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

	englishDescription := GetEnglishDescription(pokemonInformation)

	// if pokemon habitat is cave, or if it's a legendary pokemon, use yoda translation
	// if not, use shakespeare translation
	translatedDescription, err := getTranslatedDescription(pokemonInformation, englishDescription)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	getPokemonInformationResponse := &GetPokemonInformationResponse{
		Name:        name,
		Description: translatedDescription,
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

func getTranslatedDescription(pokemonInformation PokemonInformation, description string) (string, error) {
	var translatedDescription string

	params := url.Values{}
	params.Add("text", description)

	if pokemonInformation.Habitat.Name == "cave" || pokemonInformation.IsLegendary == true {
		response, err := http.PostForm("https://api.funtranslations.com/translate/yoda.json", params)
		if err != nil {
			return "", err
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return "", err
		}

		defer response.Body.Close()

		var responseObject TranslatedDescriptionResponse
		err = json.Unmarshal(responseData, &responseObject)
		if err != nil {
			return "", err
		}

		// if translation does not succeed, return original description
		if responseObject.Success.Total == 0 {
			return description, nil
		}

		translatedDescription = responseObject.Contents.Translated

	} else {
		response, err := http.PostForm("https://api.funtranslations.com/translate/shakespeare.json", params)
		if err != nil {
			return "", err
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return "", err
		}

		defer response.Body.Close()

		var responseObject TranslatedDescriptionResponse
		err = json.Unmarshal(responseData, &responseObject)
		if err != nil {
			return "", err
		}

		// if translation does not succeed, return original description
		if responseObject.Success.Total == 0 {
			return description, nil
		}

		translatedDescription = responseObject.Contents.Translated
	}
	return translatedDescription, nil
}
