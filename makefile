BINARY_NAME=main.out

build:
	go build -o ${BINARY_NAME} main.go commands.go
run:
	go build -o ${BINARY_NAME} main.go commands.go
	./${BINARY_NAME}
clean:
	go clean
	rm ${BINARY_NAME}
