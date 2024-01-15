package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/crewjam/saml/samlsp"
	"github.com/eddique/okta_go_saml/pkg/core/models"
)

func (api *ApiAdapter) InfoHandler(w http.ResponseWriter, r *http.Request) {

	email := samlsp.AttributeFromContext(r.Context(), "email")
	firstName := samlsp.AttributeFromContext(r.Context(), "firstName")
	lastName := samlsp.AttributeFromContext(r.Context(), "lastName")

	user := models.User{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
