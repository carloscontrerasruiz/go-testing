package repository

import (
	"catching-pokemons/custom_errors"
	"catching-pokemons/models"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestGetPokemonRepositorySuccess(t *testing.T) {
	assert := require.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	id := "bulbasaur"

	body, err := os.ReadFile("samples/pokeapi_response.json")
	assert.NoError(err)

	request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)
	httpmock.RegisterResponder(
		"GET",
		request,
		httpmock.NewStringResponder(200, string(body)),
	)

	pokemon, err := GetPokemonByID(id)
	assert.NoError(err)

	var expected models.PokeApiPokemonResponse
	err = json.Unmarshal([]byte(body), &expected)
	assert.NoError(err)

	assert.Equal(expected, pokemon)

}

func TestGetPokemonRepositoryNotFound(t *testing.T) {
	assert := require.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	id := "bulbasaur"

	body, err := os.ReadFile("samples/pokeapi_response.json")
	assert.NoError(err)

	request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)
	httpmock.RegisterResponder(
		"GET",
		request,
		httpmock.NewStringResponder(404, string(body)),
	)

	_, err = GetPokemonByID(id)
	assert.NotNil(err)
	assert.EqualError(custom_errors.ErrPokemonNotFound, err.Error())

}
