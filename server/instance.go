package server

import (
	"net/http"

	"github.com/marcosrgrave/go-cep-server/handlers"
)

func Run() {
	loadHandlers()

	http.ListenAndServe("localhost:8000", nil)
}

func loadHandlers() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/cep/", handlers.CEPSearch)
}
