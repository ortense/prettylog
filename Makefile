test:
	go test ./...

build:
	go build -o ./bin/prettylog cmd/prettylog/main.go

install:
	sudo cp ./bin/prettylog /usr/local/bin

demo:
	go run ./example/main.go | go run ./cmd/prettylog/main.go