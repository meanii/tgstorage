BINARY_NAME=tgstorage
BUILD_DIR=build

build:
	@echo "Building for every OS and Platform"
	@mkdir -p ${BUILD_DIR}
	GOARCH=amd64 GOOS=darwin go build -o ./${BUILD_DIR}/${BINARY_NAME}-darwin main.go
	GOARCH=arm64 GOOS=darwin go build -o ./${BUILD_DIR}/${BINARY_NAME}-darwin-arm64 main.go
	GOARCH=amd64 GOOS=linux go build -o ./${BUILD_DIR}/${BINARY_NAME}-linux main.go
	GOARCH=arm64 GOOS=linux go build -o ./${BUILD_DIR}/${BINARY_NAME}-linux-arm64 main.go
	GOARCH=amd64 GOOS=windows go build -o ./${BUILD_DIR}/${BINARY_NAME}-windows main.go

deps:
	@echo "Installing dependencies"
	@go mod tidy

update:
	@echo "Updating dependencies"
	@go get -u all
	@go mod tidy

cbuild:
	@echo "Building for current OS and Platform"
	@mkdir -p ${BUILD_DIR}
	@go build -o ./${BUILD_DIR}/${BINARY_NAME} main.go

run: cbuild
	./${BUILD_DIR}/${BINARY_NAME}

clean:
	@echo "Cleaning up"
	@rm -rf build

dev:
	@echo "Running in development mode"
	@nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go
