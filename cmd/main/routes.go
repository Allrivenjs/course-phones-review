package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func Routes() *chi.Mux {
	mux := chi.NewMux()

	//middleware
	mux.Use(
		middleware.Logger,    // log every http request
		middleware.Recoverer, // recover from panics without crashing server
	)

	//routes
	mux.Post("/smartphones", nil)
	mux.Get("/hello", helloHandler)
	return mux
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := map[string]interface{}{
		"message": "hello world",
	}
	_ = json.NewEncoder(w).Encode(res)
}
