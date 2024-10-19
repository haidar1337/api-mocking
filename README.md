# API Mocking
API Mocking is a service that allows you to create mocking endpoints to simulate API responses. It can be very useful to visualize the flow of your API endpoints requests/responses before actually implementing them.


## Get Started
Start by cloning this repository using:
```bash
git clone https://github.com/haidar1337/api-mocking.git
```

cd into the folder, then, you can change the port by navigating to `cmd/server/main.go` and changing the port of your server.
Run the server in the background using `go run .`, and navigate to `cmd/cli` to run the terminal UI using `go run .`

## API
The API for manipulating mocking endnpoints documentation can be found here [API Docs](./apidoc.md).
You are able to create a web frontend instead of a CLI and use the provided API.
