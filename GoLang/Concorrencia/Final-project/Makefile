DSN=postgres://postgres:password@localhost:5433/concurrency?sslmode=disable
BINARY_NAME=myapp.exe
REDIS=127.0.0.1:6379

## build: builds all binaries
build:
	@go build -o ${BINARY_NAME} ./cmd/web
	@echo back end built!

run: build
	@echo Starting...
	@set DSN=${DSN}&& set REDIS=${REDIS}&& .\${BINARY_NAME}
	@echo back end started!

clean:
	@echo Cleaning...
	@if exist ${BINARY_NAME} DEL ${BINARY_NAME}
	@go clean
	@echo Cleaned!

start: run

stop:
	@echo Stopping...
	@taskkill /F /IM ${BINARY_NAME} 2>NUL || exit 0
	@echo Stopped back end

restart: stop start

test:
	@echo Testing...
	go test -v ./...