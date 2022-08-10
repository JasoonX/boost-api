package handlers

import (
	"net/http"

	"github.com/markbates/goth/gothic"
)

func GetAuthProvider(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}
