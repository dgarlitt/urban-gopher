package routing

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/dgarlitt/urban-gopher/pkg/config"
	"github.com/dgarlitt/urban-gopher/pkg/log"
	"github.com/dgarlitt/urban-gopher/pkg/types"
	"github.com/dgarlitt/urban-gopher/version"
)

var out io.Writer = os.Stdout

// HealthCheck - response format for health handler
type HealthCheck struct {
	Status  string `json:"status"`
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Branch  string `json:"branch"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Info.Println(r)
	w.Header().Set("Content-Type", "application/json")
	jsonstr, _ := json.Marshal(&HealthCheck{
		"Alive and kicking!",
		version.Version,
		version.Commit,
		version.Branch,
	})

	w.Write(jsonstr)
}

// DefinitionHandlerType - Set the type of provider to use for definitions
type DefinitionHandlerType struct {
	Provider types.DefinitionProvider
}

func (dht *DefinitionHandlerType) definitionHandler(w http.ResponseWriter, r *http.Request) {
	var jsonstr []byte

	w.Header().Set("Content-Type", "application/json")

	params := &types.ProviderParams{
		Term:   r.URL.Query().Get("term"),
		APIKey: r.Header.Get("X-API-Key"),
	}

	if len(params.APIKey) == 0 {
		params.APIKey = config.APIKey
	}

	if len(params.APIKey) == 0 {
		errMsg := "No API key provided"
		http.Error(w, errMsg, http.StatusBadRequest)
		log.Error.Println(strconv.Itoa(http.StatusBadRequest), errMsg, r)
		return
	}

	r.Header.Set("X-API-Key", strings.Repeat("*", len(params.APIKey)))
	log.Info.Println(r)

	definition, err := dht.Provider.LookupDefinition(params)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		log.Error.Println(strconv.Itoa(http.StatusBadGateway), err.Error())
		return
	}

	if err == nil {
		jsonstr, _ = json.Marshal(definition)
	}

	w.Write(jsonstr)
}
