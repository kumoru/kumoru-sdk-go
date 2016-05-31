BUILD_FLAGS :=  "-s -w -X main.BuildStamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.GitVersion=`git rev-parse HEAD`"

default: clean build

linux-binary:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags $(BUILD_FLAGS) -o builds/linux/kumoru main.go

osx-binary:
	CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -ldflags $(BUILD_FLAGS) -o builds/osx/kumoru main.go

windows-binary:
	CGO_ENABLED=0 GOOS=windows go build -a -installsuffix cgo -ldflags $(BUILD_FLAGS) -o builds/windows/kumoru main.go
build:  test 
	go build -a -ldflags $(BUILD_FLAGS) -o kumoru main.go

install: build mv-bin

mv-bin:
	cp kumoru ${GOPATH}/bin/

clean:
	rm -f kumoru
	rm -f builds/osx/kumoru
	rm -f builds/linux/kumoru
	rm -f builds/windows/kumoru

test:
	go test -cover ./...

release: clean test osx-binary linux-binary windows-binary
