package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// GET endpoint
	r.Get("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		userId := chi.URLParam(r, "userId")
		w.Write([]byte("Retrieve user: " + userId))
	})

	// POST endpoint
	r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Create a new user"))
	})

	// PUT endpoint
	r.Put("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		userId := chi.URLParam(r, "userId")
		w.Write([]byte("Update user: " + userId))
	})

	// DELETE endpoint
	r.Delete("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		userId := chi.URLParam(r, "userId")
		w.Write([]byte("Delete user: " + userId))
	})

	http.ListenAndServe(":8080", r)
}
