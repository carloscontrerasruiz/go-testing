package controller

import (
	"catching-pokemons/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestGetPokemon(t *testing.T) {
	assert := require.New(t)

	r, err := http.NewRequest("GET", "/pokemon/{id}", nil)
	assert.NoError(err)

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "bulbasaur",
	}

	r = mux.SetURLVars(r, vars)

	GetPokemon(w, r)

	expectedBodyResponse, err := os.ReadFile("samples/api_response.json")
	assert.NoError(err)

	var expected models.Pokemon
	err = json.Unmarshal([]byte(expectedBodyResponse), &expected)
	assert.NoError(err)

	var actualPokemon models.Pokemon
	err = json.Unmarshal([]byte(w.Body.Bytes()), &actualPokemon)
	assert.NoError(err)

	assert.Equal(http.StatusOK, w.Code)
	assert.Equal(expected, actualPokemon)
}
