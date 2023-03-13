package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/marcosrgrave/go-cep-server/controllers"
)

func Home(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	cepParam := params.Get("cep")

	if len(cepParam) == 8 && cepParam != "" {
		viaCEP, err := controllers.CEPSearch(cepParam)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return

		}
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(viaCEP)
		return
	}

	if len(cepParam) != 8 && cepParam != "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("CEP should follow this format: 01001000\nPlease, try again."))
		return
	}

	w.WriteHeader(http.StatusOK)
	body := map[string]string{"msg": "Home Page"}
	json.NewEncoder(w).Encode(body)
}
