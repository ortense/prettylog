build:
	go build -o ./bin/prettylog cmd/main.go

install:
	sudo cp ./bin/prettylog /usr/local/bin

demo:
	go run ./example/main.go | go run ./cmd/main.go
