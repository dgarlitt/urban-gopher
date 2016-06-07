package main

import (
	builtinLog "log"
	"net/http"
	"os"

	"github.com/dgarlitt/urban-gopher/pkg/config"
	customLog "github.com/dgarlitt/urban-gopher/pkg/log"
	"github.com/dgarlitt/urban-gopher/pkg/routing"
)

func main() {
	config.APIKey = os.Getenv("URBAN_GOPHER_API_KEY")
	customLog.SetOutput()
	routing.SetupRoutes()
	builtinLog.Fatal(http.ListenAndServe(":8008", nil))
}
