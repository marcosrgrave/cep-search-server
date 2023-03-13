package handlers

import (
	"encoding/json"
	"net/http"
)

func CEPSearch(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	body := map[string]string{"msg": "CEP Page"}
	json.NewEncoder(w).Encode(body)
}
