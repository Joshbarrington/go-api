#
# Makefile variables
#
NC=\033[0m
RED=\033[1;31m
GRN=\033[1;32m
YEL=\033[1;33m
TIMESTAMP := $$(date "+%Y-%m-%d %H:%M:%S")#

#
# Go targets
#
lint:
	@printf "$(TIMESTAMP) $(YEL)Running Go formatting and linting...$(NC)\n"
	go fmt ./...
	golint ./...

run-tests:
	@printf "$(TIMESTAMP) $(YEL)Running Go tests...$(NC)\n"
	cd httpd/routes && go test

dev:
	@printf "$(TIMESTAMP) $(YEL)Running Go dev server...$(NC)\n"
	go run httpd/main.go

build:
	@printf "Project is going to be rebuilt\n"
	docker-compose build
	@printf "Build over\n"

docker-compose-up:
	@printf "$(TIMESTAMP) $(YEL)Launching Go API and MongoDB...$(NC)\n"
	-docker-compose rm -fs
	docker-compose run api sh -c 'make run-tests'
