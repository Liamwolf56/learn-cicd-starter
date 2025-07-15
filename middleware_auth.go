package main

import (
	"net/http"
)

type User struct {
	ID    int
	Email string
}

type authedHandler func(http.ResponseWriter, *http.Request, User)

func (fn authedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// This is where you'd fetch the user via API key etc.
	var user User
	fn(w, r, user)
}

func (cfg *apiConfig) middlewareAuth(next authedHandler) http.Handler {
	return next
}
