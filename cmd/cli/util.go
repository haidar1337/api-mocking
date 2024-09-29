package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func handleStatusCodeErr(code int) error {
	if code > 299 {
		return fmt.Errorf("something went wrong: %d", code)
	}
	return nil
}

func structureEndpoints(endpoints []mockendpoint) string {
	out := ""
	for idx, endpoint := range endpoints {
		out += fmt.Sprintf("%d. %s %s\n", idx+1, endpoint.Method, endpoint.Endpoint)
	}
	return out
}

func handleIDSelection(msg string, cfg *config, endpoints []mockendpoint) int {
	scanner := bufio.NewScanner(os.Stdin)
	var id int
	var err error
	for {
		fmt.Println(msg)
		fmt.Print(structureEndpoints(endpoints))
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			break
		}
		id, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("invalid input")
			continue
		}
		if id > len(cfg.commands) {
			fmt.Println("endpoint does not exist")
			continue
		}
		break
	}
	return id
}

func getEndpoints(cfg *config) ([]mockendpoint, error) {
	url := cfg.baseURL + "/mock/endpoints"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := cfg.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	err = handleStatusCodeErr(res.StatusCode)
	if err != nil {
		return nil, err
	}

	endpoints := []mockendpoint{}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &endpoints)
	if err != nil {
		return nil, err
	}

	return endpoints, nil
}

func structureResponseBody(res map[string]any) string {
	response := "Response: {\n"
	for k, v := range res {
		response += fmt.Sprintf("  \"%v\": %v\n", k, v)
	}
	response += "}"
	return response
}

func structureEndpoint(endpoint mockendpoint) string {
	out := ""

	msg := fmt.Sprintf("Endpoint Route: %s\nRequest Method: %s\nStatus Code: %d\nDelay: %d\n", endpoint.Endpoint, endpoint.Method, endpoint.Response.StatusCode, endpoint.Delay)

	fields := "Request Fields >\n"
	for i := 0; i < len(endpoint.Request.Body); i++ {
		f := endpoint.Request.Body[i]
		fields += fmt.Sprintf("Field %s Type %s Required %v\n", f.Name, f.Type, f.Required)
	}

	response := structureResponseBody(endpoint.Response.Body)

	return out + msg + fields + response
}
