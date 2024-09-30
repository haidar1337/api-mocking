package main

type field struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
}

type mockendpointrequest struct {
	Body []field `json:"request_body"`
}

type mockendpointresponse struct {
	StatusCode int            `json:"status_code"`
	Body       map[string]any `json:"response_body"`
}

type MockEndpointErrorSimulation struct {
	ErrorStatusCode int    `json:"error_status_code"`
	Body            string `json:"error_body"`
}

type mockendpoint struct {
	Endpoint        string                      `json:"endpoint"`
	Method          string                      `json:"method"`
	Delay           uint                        `json:"delay"`
	Request         mockendpointrequest         `json:"request"`
	Response        mockendpointresponse        `json:"response"`
	ErrorSimulation MockEndpointErrorSimulation `json:"error_simulation"`
}
