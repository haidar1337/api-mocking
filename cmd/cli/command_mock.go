package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMock(cfg *config, args ...string) error {
	endpoints, err := getEndpoints(cfg)
	if err != nil {
		return err
	}

	id := handleSelection("Which endpoint would you like to mock?\nChoose from the list below by typing in the number of the endpoint or type exit to exit", cfg, endpoints)
	if id == 0 {
		return nil
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/mock/%d", cfg.baseURL, id), nil)
	if err != nil {
		return err
	}
	fmt.Printf("Sending a %s request to %s...\n", endpoints[id-1].Method, endpoints[id-1].Endpoint)

	res, err := cfg.httpClient.Do(req)
	if err != nil {
		return err
	}
	mockendpoint := []interface{}{}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &mockendpoint)
	if err != nil {
		return err
	}

	fmt.Println(mockendpoint)
	return nil
}
