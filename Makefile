BINARY_NAME=bendisApp

build:
	@go mod vendor
	@echo "Building Bendis..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Bendis built!"

run: build
	@echo "Starting Bendis..."
	@./tmp/${BINARY_NAME} &
	@echo "Bendis started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

start_compose:
	@docker-compose up -d

stop_compose:
	@docker-compose down

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping Bendis..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Bendis!"

restart: stop start