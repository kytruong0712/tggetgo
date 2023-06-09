package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/api/v1/foo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!!!"))
	})
	http.ListenAndServe(":3333", r)
}
