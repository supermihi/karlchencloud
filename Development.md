# Development

## API
We use the [gRPC](https://grpc.io/) protocol that allows for type-safe APIs and real-time streaming of game events to other players.
## Server
The server is implemented in [Go](https://golang.org/). To get started,
open the repository root in the Go IDE of your choice and run `cmd/server/main.go`.

To use the API in the browser via [grpc-web](https://github.com/grpc/grpc-web), a grpc-web enabled proxy is needed. You can run
```
docker-compose up
```
in the `envoy` directory to run envoy with a configuration that is compatible with the Karlchencloud server.


## Web App
See [frontend/README.md](frontend/README.md)