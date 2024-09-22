package mockingapi

type MockEndpoint struct {
	Endpoint string `json:"endpoint"`
	Method string `json:"method"`
	Data any `json:"data"`
	Response MockEndpointResponse `json:"response"`
}

type MockEndpointResponse struct {
	Status string `json:"status"`
	StatusCode int `json:"status_code"`
	Body string `json:"body"`
}