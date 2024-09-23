package db

import "errors"

type MockEndpoint struct {
	Endpoint string `json:"endpoint"`
	Method string `json:"method"`
	Delay uint
	Request MockEndpointRequest `json:"request"`
	Response MockEndpointResponse `json:"response"`
}

type MockEndpointRequest struct {
	Body []Field `json:"request_body"`
}

type MockEndpointResponse struct {
	Message string `json:"message"`
	StatusCode int `json:"status_code"`
	Body any `json:"response_body"`
}

type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Required bool `json:"required"`
}

func (db *DB) CreateMockEndpoint(mep MockEndpoint) (MockEndpoint, error) {
	structure, err := db.loadDB()
	if err != nil {
		return MockEndpoint{}, err
	}

	id := len(structure.MockEndpoints) + 1
	structure.MockEndpoints[id] = mep
	
	err = db.writeDB(structure)
	if err != nil {
		return MockEndpoint{}, err
	}

	return mep, nil
}

func (db *DB) GetEndpoints() ([]MockEndpoint, error) {
	structure, err := db.loadDB()
	if err != nil {
		return nil, err
	}

	endpoints := make([]MockEndpoint, 0)
	for _, endpoint := range structure.MockEndpoints {
		endpoints = append(endpoints, endpoint)
	}

	return endpoints, nil
}

func (db *DB) GetEndpointById(id int) (MockEndpoint, error) {
	structure, err := db.loadDB()
	if err != nil {
		return MockEndpoint{}, err
	}

	for idx, endpoint := range structure.MockEndpoints {
		if idx == id {
			return endpoint, nil
		}
	}

	return MockEndpoint{}, errors.New("endpoint not found")
}