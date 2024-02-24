BINARY_NAME=gpu
 
build:
	go build -o ./bin/${BINARY_NAME} cmd/gpud/main.go
 
run:
	go build -o ./bin/${BINARY_NAME} cmd/gpud/main.go
	./bin/${BINARY_NAME}
 
clean:
	go clean
	rm bin/${BINARY_NAME}
