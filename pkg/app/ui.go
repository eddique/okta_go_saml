package app

import (
	"net/http"
	"os"
)

func (api *ApiAdapter) FileServer(w http.ResponseWriter, r *http.Request) {
	path := "./dist" + r.URL.Path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		http.ServeFile(w, r, "./dist/index.html")
	} else {
		http.ServeFile(w, r, path)
	}
}
