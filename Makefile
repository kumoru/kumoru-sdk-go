BUILD_FLAGS :=  "-s -w -X main.gitVersion=`git rev-parse HEAD`"
SYSTEM = `uname -s`

#SDK
.PHONY: default
default: test-sdk

.PHONY: test-sdk
test-sdk:
	@rm -f pkg/coverage/coverage.txt
	@rm -f pkg/coverage/coverage.tmp
	echo 'mode: atomic' > pkg/coverage/coverage.txt && go list ./pkg/... | xargs -n1 -I{} sh -c 'go test -covermode=atomic -coverprofile=pkg/coverage/coverage.tmp {} && tail -n +2 pkg/coverage/coverage.tmp >> pkg/coverage/coverage.txt'

.PHONY: sdk-coverage
sdk-coverage:
	go tool cover -html=./pkg/coverage/coverage.txt -o ./pkg/coverage/full-report.html

#Kumoru CLI
.PHONY: default
default: clean build

.PHONY: linux-binary
linux-binary:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags $(BUILD_FLAGS) -o client/kumoru/builds/Linux/kumoru client/kumoru/main.go

.PHONY: darwin-binary
darwin-binary:
	CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -ldflags $(BUILD_FLAGS) -o client/kumoru/builds/Darwin/kumoru client/kumoru/main.go

.PHONY: windows-binary
windows-binary:
	CGO_ENABLED=0 GOOS=windows go build -a -installsuffix cgo -ldflags $(BUILD_FLAGS) -o client/kumoru/builds/Windows/kumoru client/kumoru/main.go

.PHONY: build-cli
build-cli:  test-cli
	go build -a -ldflags $(BUILD_FLAGS) -o client/kumoru/builds/$(SYSTEM)/kumoru client/kumoru/main.go

.PHONY: install-cli
install-cli: build-cli mv-bin

.PHONY: mv-bin
mv-bin:
	cp client/kumoru/builds/$(SYSTEM)/kumoru ${GOPATH}/bin/

.PHONY: clean
clean:
	rm -f client/kumoru/kumoru
	rm -f client/kumoru/builds/Darwin/kumoru
	rm -f client/kumoru/builds/Linux/kumoru
	rm -f client/kumoru/builds/Windows/kumoru

.PHONY: test-cli
test-cli:
	@rm -f client/coverage/coverage.txt
	@rm -f client/coverage/coverage.tmp
	echo 'mode: atomic' > client/coverage/coverage.txt && go list ./client/... | xargs -n1 -I{} sh -c 'go test -covermode=atomic -coverprofile=./client/coverage/coverage.tmp {} && tail -n +2 ./client/coverage/coverage.tmp >> ./client/coverage/coverage.txt'

.PHONY: cli-coverage
cli-coverage:
	go tool cover -html=./client/coverage/coverage.txt -o ./client/coverage/full-report.html

.PHONY: release-cli
release-cli: clean test-cli darwin-binary linux-binary windows-binary
