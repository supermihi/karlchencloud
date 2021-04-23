# Development

## API
We use the [gRPC](https://grpc.io/) protocol that allows for type-safe APIs and real-time streaming of game events to other players.
## Server
The server is implemented in [Go](https://golang.org/). To get started,
open the repository root in the Go IDE of your choice and run `cmd/server/main.go`.

The (raw) grpc server is automatically wrapped with the [grpc-web Go proxy](https://github.com/improbable-eng/grpc-web/tree/master/go/grpcwebproxy) by improbable
to allow for grpc-web (browser) connections.


## Web App
See [frontend/README.md](frontend/README.md)

## Development Parameters
The following debug parameters can be specified as environment variables:

| Variable    | Example             | Explanation                                        |
|-------------|---------------------|----------------------------------------------------|
| KC_DBG_SEED | 1475759920238344204 | Overwrite for randomized events (e.g. dealt cards) |
