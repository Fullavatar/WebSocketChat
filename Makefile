CLIENT_SRC = ./client
SERVER_SRC = ./main.go

CLIENT_BIN = ./builds/Client.exe
SERVER_BIN = ./builds/Server.exe

MAKE_PATH = %CD%

.PHONY: build-docker
build-docker: ## Build Docker image
	docker build --no-cache -t websocketchat .

.PHONY: build-client
build-client: ## Create the Client.exe
	godot --headless --path $(CLIENT_SRC) --export-release "Windowsx86_64" ../$(CLIENT_BIN)

.PHONY: build-server
build-server: ## Create the Server.exe
	cd server && go build -o ../$(SERVER_BIN) $(SERVER_SRC)

.PHONY: clean-build
clean-build: ## Remove Client and Server binaries
	if exist $(CLIENT_BIN) del /F /Q $(CLIENT_BIN)
	if exist $(SERVER_BIN) del /F /Q $(SERVER_BIN)

.PHONY: run-docker
run-docker: ## Run Docker image
	docker stop websocketchat || true
	docker rm websocketchat || true
	docker run -d -p 5000:5000 --name websocketchat websocketchat

.PHONY: run-client
run-client: ## Run Client
	$(CLIENT_BIN)

.PHONY: run-server
run-server: ## Run Server in a new terminal
	cmd.exe /c start "C:\Users\loica\Desktop\ChatWS\builds\Server.exe"

.PHONY: test
test: ## Run tests
	cd .\server && go test ./...

.PHONY: deps
deps: ## Install new dependencies
	cd .\server && go mod tidy -v