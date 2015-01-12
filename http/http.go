package http

import (
	"net/http"
	"os"

	gorillaHandlers "github.com/gorilla/handlers"

	"github.com/bigsurge/go-test/handlers"
	"github.com/gorilla/mux"
)

func init() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/new", handlers.NewTokenHandler)
	r.HandleFunc("/health", handlers.HealthHandler)
	r.HandleFunc("/test/{token:[a-f0-9]}", handlers.CustomHandler)

	// Only allow exact tokens with GETs and PUTs
	r.HandleFunc("/{token:[a-f0-9]{32}}", handlers.TokenHandler).
		Methods("GET", "PUT")
	r.HandleFunc("/{token:[a-f0-9]{32}}/", handlers.TokenHandler).
		Methods("GET", "PUT")
	r.HandleFunc("/{token:[a-f0-9]{32}}/{machine}", handlers.TokenHandler).
		Methods("GET", "PUT", "DELETE")
	r.HandleFunc("/{token:[a-f0-9]{32}}/_config/size", handlers.TokenHandler).
		Methods("GET")
	
	
	r.HandleFunc("/test/q/{token:[a-f0-9]}", handlers.TokenHandler).
		Methods("GET", "PUT")
	r.HandleFunc("/test/q/{token:[a-f0-9]}/", handlers.TokenHandler).
		Methods("GET", "PUT")
	r.HandleFunc("/test/q/{token:[a-f0-9]}/{machine}", handlers.TokenHandler).
		Methods("GET", "PUT", "DELETE")
	r.HandleFunc("/test/q/{token:[a-f0-9]}/_config/size", handlers.TokenHandler).
		Methods("GET")	

	logH := gorillaHandlers.LoggingHandler(os.Stdout, r)

	http.Handle("/", logH)
}
