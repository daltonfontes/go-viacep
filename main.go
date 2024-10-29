package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Address struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

type Client struct {
	httpClient *http.Client
	baseUrl    string
}

func (c *Client) GetCep(cep string) (*Address, error) {
	url := fmt.Sprintf("%s/%s/json", c.baseUrl, cep)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var address Address
	err = json.NewDecoder(resp.Body).Decode(&address)
	if err != nil {
		return nil, err
	}

	return &address, nil
}

func main() {
	client := &Client{
		httpClient: &http.Client{},
		baseUrl:    "https://viacep.com.br/ws",
	}

	endereco, err := client.GetCep("27511300")
	if err != nil {
		fmt.Println("Erro ao buscar CEP:", err)
		return
	}

	fmt.Printf("CEP: %s\nLogradouro: %s\nBairro: %s\nCidade: %s\nUF: %s\n",
		endereco.Cep,
		endereco.Logradouro,
		endereco.Bairro,
		endereco.Localidade,
		endereco.UF)
}
