package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCep struct {
	// used this tool to convert JSON to Go struct: https://mholt.github.io/json-to-go/
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func CEPSearch(cep string) (*ViaCep, error) {
	apiURL := "http://viacep.com.br/ws/" + cep + "/json/"

	req, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var data ViaCep

	err = json.Unmarshal(res, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
