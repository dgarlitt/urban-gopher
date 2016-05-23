package routing

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgarlitt/urban-gopher/pkg/types"
	"github.com/dgarlitt/urban-gopher/version"
	"github.com/stretchr/testify/assert"
)

func TestIndexHandlerResponse(t *testing.T) {
	assert := assert.New(t)
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	hc := &HealthCheck{
		"Alive and kicking!",
		"1.2.3",
		"abc123xyz",
		"my-branch",
	}

	version.SetVersion(hc.Version)
	version.SetCommit(hc.Commit)
	version.SetBranch(hc.Branch)

	expectedResponseCode := 200
	expectedResponseHead := "application/json"
	expectedResponseBody, _ := json.Marshal(hc)

	healthHandler(response, request)

	assert.Equal(expectedResponseCode, response.Code)
	assert.Equal(expectedResponseHead, response.HeaderMap.Get("Content-Type"))
	assert.Equal(string(expectedResponseBody), response.Body.String())
}

type FakeDefinitionProvider struct{}

// LookupDefinition - test method for looking up a definition
func (provider *FakeDefinitionProvider) LookupDefinition(params *types.ProviderParams) (definition *types.Definition, err error) {
	term := params.Term
	apikey := params.APIKey
	if term == "Bob Dole" {
		definition = &types.Definition{Word: term, Text: "APIKey is " + apikey}
	} else {
		err = errors.New("Error: Term = " + term + ", APIKey = " + apikey)
	}

	return
}

func TestDictionaryHandlerResponse(t *testing.T) {
	assert := assert.New(t)
	request, _ := http.NewRequest("GET", "/define?term=Bob+Dole", nil)
	request.Header.Set("X-API-Key", "bobloblaw")
	response := httptest.NewRecorder()

	expectedResponseCode := 200
	expectedResponseHead := "application/json"
	expectedResponseBody := `{"word":"Bob Dole","definition":"APIKey is bobloblaw"}`

	DefinitionHandlerProvider = &DefinitionHandlerType{Provider: &FakeDefinitionProvider{}}
	router := SetupRoutes()
	router.ServeHTTP(response, request)

	assert.Equal(expectedResponseCode, response.Code)
	assert.Equal(expectedResponseHead, response.HeaderMap.Get("Content-Type"))
	assert.Equal(expectedResponseBody, response.Body.String())
}
