package main

import (
	"net/http"
)

type User struct {
	ID    int
	Email string
}

type authedHandler func(http.ResponseWriter, *http.Request, User)

func (cfg *apiConfig) middlewareAuth(next authedHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate user authentication; replace with real logic later
		user := User{ID: 1, Email: "user@example.com"}
		next(w, r, user)
	})
}
