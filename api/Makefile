APP_NAME=app
APP_PATH=cmd/api-server/
ENTRY=app.go
BUILD_FILE_PATH=bin/

run:
	go run $(APP_PATH)$(ENTRY)

build:
	go build -o $(BUILD_FILE_PATH)$(APP_NAME) $(APP_PATH)$(ENTRY)

test:
	go test ./...

clean:
	rm -f $(BUILD_FILE_PATH)$(APP_NAME)
