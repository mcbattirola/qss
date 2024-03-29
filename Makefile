all: build

build:
	go build -o dist/main -ldflags "-H=windowsgui" main.go

run:
	go run main.go