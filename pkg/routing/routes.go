package routing

import (
	"net/http"

	"github.com/dgarlitt/urban-gopher/pkg/remoteAPIs/urbanDictionary"
	"github.com/gorilla/mux"
)

// DefinitionHandlerProvider - The service that will provide definitions
var DefinitionHandlerProvider = &DefinitionHandlerType{Provider: &urbanDictionary.DefinitionProvider{}}

// SetupRoutes - Sets up all routes for the service
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", healthHandler)

	router.HandleFunc("/definition", DefinitionHandlerProvider.definitionHandler)

	http.Handle("/", router)

	return router
}
