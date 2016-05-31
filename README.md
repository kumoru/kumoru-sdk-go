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

* You can download the latest release from [Releases](https://github.com/kumoru/kumoru-cli/releases).
* See [README.md](client/kumoru/README.md) for specific instructions on bulding the CLI locally.

## Contributing

1. Fork this repo
1. Make your changes
1. Submit a Pull Request

## Authors

* Victor Palma <victor@kumoru.io>
* Ryan Richard <ryan@kumoru.io>
