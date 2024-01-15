package main

import (
	"log"
	"net/http"

	"github.com/eddique/okta_go_saml/pkg/app"
	"github.com/eddique/okta_go_saml/pkg/core/middleware"
	"github.com/gorilla/mux"
)

func main() {
	api := app.NewAPIAdapter()
	samlSP, err := middleware.SamlMiddleware()
	if err != nil {
		log.Fatalln(err)
	}

	router := mux.NewRouter()

	router.PathPrefix("/api/v1/whoami").HandlerFunc(api.InfoHandler)
	router.PathPrefix("/").Handler(http.HandlerFunc(api.FileServer))

	app := samlSP.RequireAccount(router)

	http.Handle("/", samlSP.RequireAccount(app))
	http.Handle("/saml/", samlSP)
	http.ListenAndServe(":3000", nil)
}
