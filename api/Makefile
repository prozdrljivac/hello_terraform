APP_NAME=app
ENTRY=app.go
BUILD_FILE_PATH=bin/

run:
	go run $(ENTRY)

build:
	go build -o ${BUILD_FILE_PATH}$(APP_NAME) $(ENTRY)

test:
	go test ./...

clean:
	rm -f ${BUILD_FILE_PATH}$(APP_NAME)
