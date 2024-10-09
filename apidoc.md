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

### DELETE /mock/endpoints/{endpointId}
Delets a mock endpoint.

Response Body: 204 No Content
