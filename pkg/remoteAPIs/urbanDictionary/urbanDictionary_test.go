package urbanDictionary

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgarlitt/urban-gopher/pkg/types"
	"github.com/stretchr/testify/suite"
)

type UrbanDictionaryTestSuite struct {
	suite.Suite
	FullResponse  string
	EmptyResponse string
}

func (suite *UrbanDictionaryTestSuite) SetupTest() {
	suite.FullResponse = `{"tags":["tag1","tag2","tag3"],"result_type":"exact","list":[{"defid":12345,"word":"fake1","author":"authorone","permalink":"http://fakepermalink.com","definition":"Definition number one.","example":"example text one","thumbs_up":3000,"thumbs_down":2000,"current_vote":""},{"defid":67890,"word":"fake2","author":"authortwo","permalink":"http://fakepermalink.com","definition":"Definition number two.","example":"example text two","thumbs_up":300,"thumbs_down":250,"current_vote":""},{"defid":13579,"word":"fake3","author":"authorthree","permalink":"http://fakepermalink.com","definition":"Definition number three.","example":"example text three","thumbs_up":0,"thumbs_down":0,"current_vote":""}],"sounds":["http://media.urbandictionary.com/sound/fake-111.mp3","http://wav.urbandictionary.com/fake-222.wav"]}`
	suite.EmptyResponse = `{"tags":[],"result_type":"no_results","list":[],"sounds":[]}`
}

func (suite *UrbanDictionaryTestSuite) TestLookupDefinitionFull() {
	expectedWord := "fake2"
	expectedText := "Definition number two."
	definition, err := makeRequest(suite.FullResponse)

	suite.Equal(nil, err)
	suite.Equal(expectedWord, definition.Word)
	suite.Equal(expectedText, definition.Text)
}

func (suite *UrbanDictionaryTestSuite) TestLookupDefinitionEmpty() {
	definition, err := makeRequest(suite.EmptyResponse)

	suite.Equal("No results found for term: \"fake\".", err.Error())
	suite.Equal("", definition.Word)
	suite.Equal("", definition.Text)
}

func (suite *UrbanDictionaryTestSuite) TestLookupDefintionError() {
	definition, err := makeRequest("error")

	suite.NotNil(err)
	suite.Equal("", definition.Word)
	suite.Equal("", definition.Text)
}

func makeRequest(response string) (definition *types.Definition, err error) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if response == "error" {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, response)
		}
	}))

	if response == "error" {
		ts.Listener.Close()
	} else {
		defer ts.Close()
	}

	defURL = ts.URL
	params := &types.ProviderParams{Term: "fake", APIKey: "fakeAPIKey"}

	provider := &DefinitionProvider{}

	definition, err = provider.LookupDefinition(params)

	return
}

func TestUrbanDictionaryTestSuite(t *testing.T) {
	suite.Run(t, new(UrbanDictionaryTestSuite))
}
