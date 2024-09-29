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

	id := handleIDSelection("Which endpoint would you like to mock?\nChoose from the list below by typing in the number of the endpoint or type exit to exit", cfg, endpoints)
	if id == 0 {
		return nil
	}

	fmt.Printf("Sending a %s request to %s...\n", endpoints[id-1].Method, endpoints[id-1].Endpoint)

	mockendpoint, err := mockEndpoint(id, cfg)
	if err != nil {
		return err
	}

	fmt.Printf("Status: %v, Time: %vs\n", mockendpoint.Response.StatusCode, mockendpoint.Delay/1000)
	response := structureResponseBody(mockendpoint.Response.Body)
	fmt.Println(response)

	return nil
}

func mockEndpoint(id int, cfg *config) (mockendpoint, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/mock/%d", cfg.baseURL, id), nil)
	if err != nil {
		return mockendpoint{}, err
	}

	res, err := cfg.httpClient.Do(req)
	if err != nil {
		return mockendpoint{}, err
	}
	mockendpoint := mockendpoint{}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return mockendpoint, err
	}
	err = json.Unmarshal(body, &mockendpoint)
	if err != nil {
		return mockendpoint, err
	}

	return mockendpoint, nil
}
