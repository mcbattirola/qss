BUILD_DIST=dist
BIN_NAME=qss
BUILD_PATH=$(BUILD_DIST)/$(BIN_NAME)

all: build

build:
	go build -o $(BUILD_PATH) main.go

# build on Windows
build-w:
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_PATH).exe -ldflags "-H=windowsgui" main.go

run:
	go run main.go

test:
	go test ./...