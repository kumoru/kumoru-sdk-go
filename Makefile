BUILD_FLAGS :=  "-s -w -X main.gitVersion=`git rev-parse HEAD`"
SYSTEM = `uname -s`

#SDK
default: test-sdk

test-sdk:
	go test -cover ./pkg/...

#Kumoru CLI
default: clean build

linux-binary:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags $(BUILD_FLAGS) -o client/kumoru/builds/Linux/kumoru client/kumoru/main.go

darwin-binary:
	CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -ldflags $(BUILD_FLAGS) -o client/kumoru/builds/Darwin/kumoru client/kumoru/main.go

windows-binary:
	CGO_ENABLED=0 GOOS=windows go build -a -installsuffix cgo -ldflags $(BUILD_FLAGS) -o client/kumoru/builds/Windows/kumoru client/kumoru/main.go

build-cli:  test-cli
	go build -a -ldflags $(BUILD_FLAGS) -o client/kumoru/builds/$(SYSTEM)/kumoru client/kumoru/main.go

install-cli: build-cli mv-bin

mv-bin:
	cp client/kumoru/builds/$(SYSTEM)/kumoru ${GOPATH}/bin/

clean:
	rm -f client/kumoru/kumoru
	rm -f client/kumoru/builds/Darwin/kumoru
	rm -f client/kumoru/builds/Linux/kumoru
	rm -f client/kumoru/builds/Windows/kumoru

test-cli:
	go test -cover ./client/kumoru/...

release-cli: clean test-cli darwin-binary linux-binary windows-binary
