# Kumoru SDK for Golang

This repository holds the [Kumoru.io](https://kumoru.io) SDK for golang and the official CLI.

## Installing

```shell
go get -u github.com/kumoru/kumoru-sdk-go
```

### Requirements

* go 1.6

### The SDK

Each component of the SDK can be independently imported directly into your application:

```go
…
import "github.com/kumoru/kumoru-sdk-go/pkg/service/application/application.go
…
```

### The CLI

* You can download the latest release from [Releases](https://github.com/kumoru/kumoru-sdk-go/releases).
* See [README.md](client/kumoru/README.md) for specific instructions on bulding the CLI locally.

### Building locally

If you prefer to get the latest code and build it your self, you'll need to:

1. Clone this repository
2. Run `make install-cli` to build it for your local system. This will place the binary in your `${GOPATH}/bin` directory.


### Testing

The SDK and CLI can be tested independently via make:

```bash
make test-sdki
```

```bash
make test-cli
```

## Contributing

1. Fork this repo
1. Make your changes
1. Submit a Pull Request

## Authors

* Victor Palma <victor@kumoru.io>
* Ryan Richard <ryan@kumoru.io>
