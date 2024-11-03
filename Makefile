build: ## Build the application
	@echo "Generating sqlc files and building the application..."
	@find ./ -name "sqlc.yaml" -exec sqlc generate -f {} \; && weaver generate ./... && go build -o ./bin/app ./cmd/app/main.go

run: ## Run the application
	@weaver single deploy weaver.toml

run-multi: ## Run the application multi process
	@weaver multi deploy weaver.toml

docker-run: ## Run the application in docker
	@docker compose -f ./deployments/compose.yml up -d --build