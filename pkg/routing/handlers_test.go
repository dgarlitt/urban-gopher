package routing

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgarlitt/urban-gopher/pkg/types"
	"github.com/dgarlitt/urban-gopher/version"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
)

var router *mux.Router

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

type HandlersTestSuite struct {
	suite.Suite
}

func (suite *HandlersTestSuite) TestIndexHandlerResponse() {
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

	suite.Equal(expectedResponseCode, response.Code)
	suite.Equal(expectedResponseHead, response.HeaderMap.Get("Content-Type"))
	suite.Equal(string(expectedResponseBody), response.Body.String())
}

func (suite *HandlersTestSuite) TestDictionaryHandlerResponse() {
	request, _ := http.NewRequest("GET", "/define?term=Bob+Dole", nil)
	request.Header.Set("X-API-Key", "bobloblaw")
	response := httptest.NewRecorder()

	expectedResponseCode := http.StatusOK
	expectedResponseHead := "application/json"
	expectedResponseBody := fmt.Sprintf(
		`{"word":"%s","definition":"%s"}`,
		"Bob Dole",
		"APIKey is bobloblaw")

	router.ServeHTTP(response, request)

	suite.Equal(expectedResponseCode, response.Code)
	suite.Equal(expectedResponseHead, response.HeaderMap.Get("Content-Type"))
	suite.Equal(expectedResponseBody, response.Body.String())
}

func (suite *HandlersTestSuite) TestDictionaryHandlerResponseWithoutEnvAPIKey() {
	request, _ := http.NewRequest("GET", "/define?term=Bob+Dole", nil)
	response := httptest.NewRecorder()

	expectedResponseCode := http.StatusBadRequest

	router.ServeHTTP(response, request)

	suite.Equal(expectedResponseCode, response.Code)
}

func TestHandlersTestSuite(t *testing.T) {
	DefinitionHandlerProvider = &DefinitionHandlerType{Provider: &FakeDefinitionProvider{}}
	router = SetupRoutes()
	suite.Run(t, new(HandlersTestSuite))
}
