package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Start(port int) error {
	router := mux.NewRouter()
	attachHandlers(router)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	return s.ListenAndServe()
}

func attachHandlers(mux *mux.Router) {
	mux.HandleFunc("/", Hello)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}