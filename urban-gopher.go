package main

import (
	builtinLog "log"
	"net/http"

	customLog "github.com/dgarlitt/urban-gopher/pkg/log"
	"github.com/dgarlitt/urban-gopher/pkg/routing"
)

func main() {
	customLog.SetOutput()
	routing.SetupRoutes()
	builtinLog.Fatal(http.ListenAndServe(":8008", nil))
}
