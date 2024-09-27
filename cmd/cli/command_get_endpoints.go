package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func commandGet(cfg *config, args ...string) error {
	url := cfg.baseURL + "/mock/endpoints"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	res, err := cfg.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if res.StatusCode > 299 {
		fmt.Println(errors.New(fmt.Sprintf("something went wrong %d", res.StatusCode)))
		return nil
	}

	endpoints := []mockendpoint{}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(body, &endpoints)

	fmt.Println(structureResponse(endpoints))

	return nil
}

func structureResponse(endpoints []mockendpoint) string {
	out := ""
	for idx, endpoint := range endpoints {
		out += fmt.Sprintf("%d. %s %s\n", idx+1, endpoint.Method, endpoint.Endpoint)
	}
	return out
}
