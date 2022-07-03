ENV_VARS = MYSQL_URL=127.0.0.1:3306 MYSQL_USER=user MYSQL_PASS=pass MYSQL_SCHEMA=example

.PHONY: run
run: ; $(info running api...) @
	@$(ENV_VARS) go run ./application/main.go

.PHONY: containers
containers: ; $(info running containers...) @
	@$(ENV_VARS) docker-compose up -d