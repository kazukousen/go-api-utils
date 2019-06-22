package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kazukousen/go-api-utils/httputils"
)

func main() {
	r := chi.NewRouter()
	r.Use(httputils.Middlewares...)

	r.Get("/echo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	})

	httputils.Serve(r)
}
