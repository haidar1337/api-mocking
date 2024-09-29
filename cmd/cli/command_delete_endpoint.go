package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandDelete(cfg *config, args ...string) error {
	req, err := http.NewRequest("GET", cfg.baseURL+"/mock/endpoints", nil)
	if err != nil {
		return err
	}

	res, err := cfg.httpClient.Do(req)
	if err != nil {
		return err
	}
	err = handleStatusCodeErr(res.StatusCode)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	mockendpoints := []mockendpoint{}
	err = json.Unmarshal(body, &mockendpoints)
	if err != nil {
		return err
	}

	id := handleIDSelection("Which endpoint would you like to delete?\nChoose from the list below by typing in the number of the endpoint or type exit to exit", cfg, mockendpoints)
	if id == 0 {
		return nil
	}
	req, err = http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/mock/endpoints/%d", cfg.baseURL, id), nil)
	if err != nil {
		return err
	}

	res, err = cfg.httpClient.Do(req)
	if err != nil {
		return err
	}
	err = handleStatusCodeErr(res.StatusCode)
	if err != nil {
		return err
	}

	fmt.Println("endpoint deleted")
	return nil
}
