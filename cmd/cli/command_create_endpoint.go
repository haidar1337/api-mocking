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
	endpoint, err := handleEndpointCreation()
	if err != nil {
		return err
	}

	reqBody, err := json.Marshal(&endpoint)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/mock/endpoints", cfg.baseURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	res, err := cfg.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	err = handleStatusCodeErr(res.StatusCode)
	if err != nil {
		return err
	}
	ep := mockendpoint{}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &ep)
	if err != nil {
		return err
	}

	fmt.Println(structureEndpoint(endpoint))

	return nil
}

func handleEndpointCreation() (mockendpoint, error) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter endpoint request route (e.g. /api/users) > ")
	scanner.Scan()
	route := scanner.Text()
	if route[0] != '/' {
		return mockendpoint{}, errors.New(fmt.Sprintf("Route %s does not start with '/', plesae provide a valid route", route))
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
		return mockendpoint{}, errors.New(fmt.Sprintf("invalid method type, method must be of %v", supportedMethods))
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
		fmt.Print("Add request body fields? [Y/N]")
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
				"list",
			}
			if !slices.Contains(supportedTypes, sanitizedFieldtype) {
				return mockendpoint{}, errors.New(fmt.Sprintf("Field type %s is not supported, please provide one of the supported field types"))
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

	responses := make([]any, 0)
	if len(fields) > 0 {
		for i := 0; i < len(fields); i++ {
			fmt.Printf("Enter response body for field %s > ", fields[i])
			scanner.Scan()
			res := scanner.Text()

			responses = append(responses, res)
		}
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
			Body:       responses,
		},
	}
	fmt.Println(structureEndpoint(mockendpoint))

	return mockendpoint, nil
}

func structureEndpoint(endpoint mockendpoint) string {
	out := ""

	msg := fmt.Sprintf("Endpoint Route: %s\nRequest Method: %s\nStatus Code: %d\nDelay: %d\n", endpoint.Endpoint, endpoint.Method, endpoint.Response.StatusCode, endpoint.Delay)

	fields := "Request Fields >\n"
	for i := 0; i < len(endpoint.Request.Body); i++ {
		f := endpoint.Request.Body[i]
		fields += fmt.Sprintf("Field %s Type %s Required %t", f.Name, f.Type, f.Required)
	}

	response := fmt.Sprintf("Response >\n%v", endpoint.Response.Body)

	return out + msg + fields + response
}
