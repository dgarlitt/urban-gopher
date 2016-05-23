package routing

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

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
	jsonstr, err := json.Marshal(&HealthCheck{
		"Alive and kicking!",
		version.Version,
		version.Commit,
		version.Branch,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

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

	r.Header.Set("X-API-Key", strings.Repeat("*", len(params.APIKey)))
	log.Info.Println(r)

	definition, err := dht.Provider.LookupDefinition(params)

	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(err.Error()))
		return
	}

	if err == nil {
		jsonstr, err = json.Marshal(definition)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(jsonstr)
}
