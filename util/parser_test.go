package util

import (
	"catching-pokemons/models"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParserPokemonSuccess(t *testing.T) {
	c := require.New(t)

	body, err := os.ReadFile("samples/pokeapi_response.json")
	c.NoError(err)

	var response models.PokeApiPokemonResponse

	err = json.Unmarshal([]byte(body), &response)
	c.NoError(err)

	parsedPokemon, err := ParsePokemon(response)
	c.NoError(err)

	body, err = os.ReadFile("samples/api_response.json")
	c.NoError(err)

	var expectedPokemon models.Pokemon

	err = json.Unmarshal([]byte(body), &expectedPokemon)
	c.NoError(err)

	c.Equal(expectedPokemon, parsedPokemon)
}

func TestParserPokemonTypeNotFound(t *testing.T) {
	assert := require.New(t)

	body, err := os.ReadFile("samples/pokeapi_response.json")
	assert.NoError(err)

	var response models.PokeApiPokemonResponse

	err = json.Unmarshal([]byte(body), &response)
	assert.NoError(err)

	response.PokemonType = []models.PokemonType{}

	_, err = ParsePokemon(response)

	assert.NotNil(err)
	assert.EqualError(ErrNotFoundPokemonType, err.Error())
}
