package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"
	"strconv"
	"strings"
)

func commandCreate(cfg *config, args ...string) error {
	userCreatedEndpoint, err := handleEndpointCreation()
	if err != nil {
		return err
	}

	endpoint, err := createEndpoint(userCreatedEndpoint, cfg)
	if err != nil {
		return err
	}

	fmt.Println("Successfully created a mock endpoint with the following details:")
	fmt.Println(structureEndpoint(endpoint))

	return nil
}

func createEndpoint(endpoint mockendpoint, cfg *config) (mockendpoint, error) {
	reqBody, err := json.Marshal(&endpoint)
	if err != nil {
		return mockendpoint{}, err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/mock/endpoints", cfg.baseURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return mockendpoint{}, err
	}

	res, err := cfg.httpClient.Do(req)
	if err != nil {
		return mockendpoint{}, err
	}
	defer res.Body.Close()
	err = handleStatusCodeErr(res.StatusCode)
	if err != nil {
		return mockendpoint{}, err
	}
	ep := mockendpoint{}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return mockendpoint{}, err
	}
	err = json.Unmarshal(body, &ep)
	if err != nil {
		return mockendpoint{}, err
	}

	return ep, nil
}

func handleEndpointCreation() (mockendpoint, error) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter endpoint request route (e.g. /api/users) > ")
	scanner.Scan()
	route := scanner.Text()
	if route[0] != '/' {
		return mockendpoint{}, fmt.Errorf("route %s does not start with '/', plesae provide a valid route", route)
	}

	fmt.Print("Enter endpoint request method (e.g. GET) > ")
	scanner.Scan()
	method := scanner.Text()
	supportedMethods := []string{
		"GET",
		"POST",
		"DELETE",
		"PUT",
	}
	if !slices.Contains(supportedMethods, strings.ToUpper(method)) {
		return mockendpoint{}, fmt.Errorf("invalid method type, method must be of %v", supportedMethods)
	}

	fmt.Print("Enter endpoint response delay in milliseconds (e.g. 1000 for 1s) or type 0 for no delay > ")
	scanner.Scan()
	enteredDelay := scanner.Text()
	delay, err := strconv.Atoi(enteredDelay)
	if err != nil {
		return mockendpoint{}, errors.New("invalid delay")
	}
	if delay < 0 {
		return mockendpoint{}, errors.New("delay must be a positive integer")
	}

	fields := make([]field, 0)
	for {
		fmt.Print("Add request body fields? [Y/N] > ")
		scanner.Scan()
		wantsToAdd := scanner.Text()
		if strings.ToLower(wantsToAdd) == "n" {
			break
		}
		if strings.ToLower(wantsToAdd) == "y" {
			fmt.Print("Enter request body field name (e.g. email) > ")
			scanner.Scan()
			fieldName := scanner.Text()

			fmt.Print("Enter request body field type (e.g. string) > ")
			scanner.Scan()
			fieldType := scanner.Text()
			sanitizedFieldtype := strings.ToLower(fieldType)
			supportedTypes := []string{
				"string",
				"int",
				"bool",
				"float",
			}
			if !slices.Contains(supportedTypes, sanitizedFieldtype) {
				return mockendpoint{}, fmt.Errorf("field type %s is not supported, please provide one of the supported field types", sanitizedFieldtype)
			}

			fmt.Print("Is this field required? [Y/N] > ")
			scanner.Scan()
			required := scanner.Text()
			isRequired := false
			if strings.ToLower(required) == "y" {
				isRequired = true
			}

			f := field{
				Name:     fieldName,
				Type:     fieldType,
				Required: isRequired,
			}

			fields = append(fields, f)
			continue
		}
	}

	fmt.Print("Enter response status code (e.g. 200) > ")
	scanner.Scan()
	enteredCode := scanner.Text()
	statusCode, err := strconv.Atoi(enteredCode)
	if err != nil {
		return mockendpoint{}, errors.New("invalid status code")
	}
	if statusCode < 100 || statusCode > 599 {
		return mockendpoint{}, errors.New("status code must be between 100 and 599")
	}

	r := make(map[string]any, 0)
	for {
		fmt.Print("Add a response body? [Y/N] > ")
		scanner.Scan()
		response := scanner.Text()
		if strings.ToLower(response) == "n" {
			break
		}

		fmt.Print("Enter response body key (e.g. username) > ")
		scanner.Scan()
		k := scanner.Text()
		fmt.Printf("Enter %v value (e.g. haidar1337) > ", k)
		scanner.Scan()
		v := scanner.Text()

		r[k] = v
	}

	fmt.Printf("Add error simulation? [Y/N] > ")
	scanner.Scan()
	response := scanner.Text()
	var errorSimulation MockEndpointErrorSimulation
	if strings.ToLower(response) == "y" {
		fmt.Print("Enter error status code (e.g. 500) > ")
		scanner.Scan()
		userEnteredCode := scanner.Text()
		errorStatusCode, err := strconv.Atoi(userEnteredCode)
		if err != nil {
			return mockendpoint{}, err
		}
		if errorStatusCode < 400 || errorStatusCode > 599 {
			return mockendpoint{}, fmt.Errorf("invalid error status code %d, error status code must be between the range of 400 to 599", errorStatusCode)
		}

		fmt.Printf("Enter error response body (e.g. user not found) > ")
		scanner.Scan()
		errBody := scanner.Text()

		errorSimulation.Body = errBody
		errorSimulation.ErrorStatusCode = errorStatusCode
	}

	mockendpoint := mockendpoint{
		Endpoint: route,
		Method:   method,
		Delay:    uint(delay),
		Request: mockendpointrequest{
			Body: fields,
		},
		Response: mockendpointresponse{
			StatusCode: statusCode,
			Body:       r,
		},
		ErrorSimulation: errorSimulation,
	}

	return mockendpoint, nil
}
