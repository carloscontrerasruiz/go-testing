package controller

import (
	"catching-pokemons/custom_errors"
	"catching-pokemons/repository"
	"catching-pokemons/util"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_, err = w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	apiPokemon, err := repository.GetPokemonByID(id)

	if errors.Is(err, custom_errors.ErrPokemonNotFound) {
		respondwithJSON(w, http.StatusBadRequest, fmt.Sprintf("Pokemon not found: %s", err.Error()))
	}

	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, fmt.Sprintf("error while calling poke api: %s", err.Error()))
	}

	parsedPokemon, err := util.ParsePokemon(apiPokemon)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, fmt.Sprintf("error found: %s", err.Error()))
	}

	respondwithJSON(w, http.StatusOK, parsedPokemon)
}
