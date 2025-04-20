APP_NAME=app
APP_PATH=cmd/api-server/
ENTRY=app.go
BUILD_FILE_PATH=bin/

run: swag
	go run $(APP_PATH)$(ENTRY)

dev:
	go run $(APP_PATH)$(ENTRY)

build:
	go build -o $(BUILD_FILE_PATH)$(APP_NAME) $(APP_PATH)$(ENTRY)

test:
	go test ./...

clean:
	rm -f $(BUILD_FILE_PATH)$(APP_NAME)

swag:
	go install github.com/swaggo/swag/cmd/swag@v1.7.9
	swag init --generalInfo cmd/api-server/app.go --output ./docs --parseDependency