# Mock API
This file is to provide documentation about the mock API, aka the server running in the background to serve the requests coming from the terminal UI.

## Mock Resource
[//]: <> (Continue...)
```json
{
  "endpoint": "/api/login",
  "method": "POST",
  "delay": 1000,
  "request": {
    "request_body": [
      {
        "name": "email",
        "type": "string",
        "required": true
      },
      {
        "name": "password",
        "type": "string",
        "required": true
      },
    ]
  },
  "response": {
    "status_code": 201,
    "response_body": {
      "id": "jfsaof29v5n91vm3jr3q90rjq09gh94",
      "email": "test@example.com"
    }
  },
  "error_simulation": {
    "error_status_code": 400,
    "error_body": "bad request"
  }
}
```

### POST /mock/endpoints
Create a new mock endpoint.

Request Body:
```json
{
  "endpoint": "/api/login",
  "method": "POST",
  "delay": 1000,
  "request": {
    "request_body": [
      {
        "name": "email",
        "type": "string",
        "required": true
      },
      {
        "name": "password",
        "type": "string",
        "required": true
      },
    ]
  },
  "response": {
    "status_code": 201,
    "response_body": {
      "id": "jfsaof29v5n91vm3jr3q90rjq09gh94",
      "email": "test@example.com"
    }
  },
  "error_simulation": {
    "error_status_code": 400,
    "error_body": "bad request"
  }
}
```

Response Body: 201 with the same request body.

### GET /mock/endpoints/{endpointId}
Get all mock endpoints.

Response Body:
```json
{
  "endpoint": "/api/users/{userId}",
  "method": "GET",
  "delay": 2000,
  "request": {
    "request_body": [
      {
        "name": "id",
        "type": "string",
        "required": true
      }
    ]
  },
  "response": {
    "status_code": 200,
    "response_body": {
      "id": "jfsaof29v5n91vm3jr3q90rjq09gh94",
      "username": "haidar1337"
    }
  },
  "error_simulation": {
    "error_status_code": 404,
    "error_body": "user not found"
  }
},
...
```


### GET /mock/endpoints/{endpointId}
Get a mock endpoint by its id.

Response Body:
```json
{
  "endpoint": "/api/users/{userId}",
  "method": "GET",
  "delay": 2000,
  "request": {
    "request_body": [
      {
        "name": "id",
        "type": "string",
        "required": true
      }
    ]
  },
  "response": {
    "status_code": 200,
    "response_body": {
      "id": "jfsaof29v5n91vm3jr3q90rjq09gh94",
      "username": "haidar1337"
    }
  },
  "error_simulation": {
    "error_status_code": 404,
    "error_body": "user not found"
  }
}
```

### DELETE /mock/endpoints/{endpointId}
Deletes a mock endpoint.

Response Body: 204 No Content
